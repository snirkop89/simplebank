package gapi

import (
	"context"
	"database/sql"

	db "github.com/snirkop89/simplebank/db/sqlc"
	"github.com/snirkop89/simplebank/pb"
	"github.com/snirkop89/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "unexpected database error")
	}

	err = util.CheckPassowrd(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to create access token")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuarion,
	)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to create refresh token")
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "", // TODO
		ClientIp:     "", // TODO
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to create session")
	}

	rsp := pb.LoginUserResponse{
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		User:                  convertUser(user),
	}
	return &rsp, nil
}

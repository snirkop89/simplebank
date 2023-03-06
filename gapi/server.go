package gapi

import (
	"fmt"

	db "github.com/snirkop89/simplebank/db/sqlc"
	"github.com/snirkop89/simplebank/pb"
	"github.com/snirkop89/simplebank/token"
	"github.com/snirkop89/simplebank/util"
	"github.com/snirkop89/simplebank/worker"
)

// Server serves gRPC rewqurest for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config         util.Config
	store          db.Store
	tokenMaker     token.Maker
	taskDistibutor worker.TaskDistributor
}

// NewServer create a new server and setup routing
func NewServer(config util.Config, store db.Store, taskDistibutor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:         config,
		store:          store,
		tokenMaker:     tokenMaker,
		taskDistibutor: taskDistibutor,
	}

	return server, nil
}

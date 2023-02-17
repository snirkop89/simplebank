package token

import "time"

// Maker is an interface for managing token
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

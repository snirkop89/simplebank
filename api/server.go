package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/snirkop89/simplebank/db/sqlc"
)

// Server serves HTTP rewqurest for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create a new server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specified address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

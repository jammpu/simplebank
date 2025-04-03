package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jammpu/simplebank/db/sqlc"
)

// Server servirá todas las peticiones http
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer crea una nueva HTTP server y configura las rutas
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start corre el HTTP server con una direccion especifica
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

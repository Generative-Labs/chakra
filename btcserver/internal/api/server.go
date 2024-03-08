package api

import (
	"fmt"

	"github.com/generativelabs/btcserver/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	backend *db.Backend
	engine  *gin.Engine
}

func New(backend *db.Backend) *Server {
	server := &Server{
		backend: backend,
	}

	r := gin.Default()
	r.Use(CORSMiddleware())
	SetupRoutes(r, server)

	server.engine = r
	return server
}

func (s Server) Run(servicePort int) error {
	return s.engine.Run(fmt.Sprintf(":%d", servicePort))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

package api

import (
	"context"
	"fmt"
	"time"

	"github.com/NethermindEth/starknet.go/account"
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/generativelabs/btcserver/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Ctx               context.Context
	backend           *db.Backend
	engine            *gin.Engine
	ChakraAccount     *account.Account
	ContractAddress   string
	ScheduleTimeWheel time.Time
	btcClient         *btc.Client
}

func NewServer(ctx context.Context, backend *db.Backend, chakraAccount *account.Account,
	contractAddress string, btcClient *btc.Client,
) *Server {
	server := &Server{
		Ctx:             ctx,
		backend:         backend,
		ChakraAccount:   chakraAccount,
		ContractAddress: contractAddress,
		btcClient:       btcClient,
	}

	go server.TimeWheelSchedule()
	go server.UpdateStakeStatus()

	r := gin.Default()
	r.Use(CORSMiddleware())
	SetupRoutes(r, server)

	server.engine = r
	return server
}

func (s *Server) Run(servicePort int) error {
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

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

	r := gin.Default()
	r.Use(CORSMiddleware())
	SetupRoutes(r, server)

	server.engine = r
	return server
}

func (s *Server) Run(servicePort int) error {
	go s.TimeWheelSchedule()
	go s.UpdateStakeFinalizedStatus()

	return s.engine.Run(fmt.Sprintf(":%d", servicePort))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(204)
		//	return
		//}

		c.Next()
	}
}

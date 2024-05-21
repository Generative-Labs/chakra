package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, server *Server) {
	r.GET("/api/stakes_list", server.GetStakeListByStaker)

	r.POST("/api/stake_btc", server.SubmitProofOfStake)

	r.GET("/api/stake_index", server.GetStakeIndexByStaker)
}

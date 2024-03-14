package api

import (
	"net/http"

	"github.com/generativelabs/btcserver/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HTTPResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func JSONResp(c *gin.Context, code int, res *HTTPResponse) {
	c.JSON(code, res)
}

// GetStakeListByStaker Get staking history
func (s *Server) GetStakeListByStaker(c *gin.Context) {
	respData := &HTTPResponse{Msg: "Ok"}

	var staker types.StakerReq
	if err := c.BindQuery(&staker); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	total, err := s.backend.QueryStakesCountByStaker(staker.Staker)
	if err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusInternalServerError, respData)
		return
	}

	stakes, err := s.backend.QueryStakesByStaker(staker.Staker, staker.Page, staker.Size)
	if err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusInternalServerError, respData)
		return
	}

	srakeList := make([]*types.StakeInfoResp, 0)
	for _, s := range stakes {
		srakeList = append(srakeList, &types.StakeInfoResp{
			Staker:         s.Staker,
			Tx:             s.Tx,
			Start:          s.Start,
			Durnation:      s.Duration,
			Amount:         s.Amount,
			RewardReceiver: s.RewardReceiver,
		})
	}

	respData.Data = map[string]interface{}{
		"total_count": total,
		"data_list":   srakeList,
	}
	JSONResp(c, http.StatusOK, respData)
}

// SubmitProofOfStake Submit proof of Stake
func (s *Server) SubmitProofOfStake(c *gin.Context) {
	respData := &HTTPResponse{Msg: "Ok"}

	var stakeInfo types.StakeInfoReq
	if err := c.Bind(&stakeInfo); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	// verify reward receiver signature
	if err := s.btcClient.CheckRewardAddressSignature(stakeInfo.StakerPublicKey, stakeInfo.RewardReceiver,
		stakeInfo.ReceiverSignature, int32(stakeInfo.Timestamp),
	); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	log.Info().Msgf("[HTTP] SubmitProofOfStake: %+v ", stakeInfo)
	// storage
	err := s.backend.CreateStake(stakeInfo.Staker,
		stakeInfo.StakerPublicKey,
		stakeInfo.TxID,
		stakeInfo.Start,
		stakeInfo.Duration,
		stakeInfo.Amount,
		stakeInfo.RewardReceiver,
		stakeInfo.ReceiverSignature,
		stakeInfo.Timestamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, respData)
}

package api

import (
	"net/http"
	"time"

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
			Duration:       s.Duration,
			Deadline:       s.Deadline,
			Amount:         s.Amount,
			RewardReceiver: s.RewardReceiver,
			Reward:         s.Reward,
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

	// verify reward
	// reward := utils.CalculateReward(float64(stakeInfo.Amount), float64(stakeInfo.Duration))
	// if reward != float64(stakeInfo.Reward) {
	// 	respData.Msg = errors.New("reward calculation error").Error()
	// 	JSONResp(c, http.StatusBadRequest, respData)
	// 	return
	// }

	if err := s.btcClient.CheckTxID(stakeInfo.TxID); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	// verify reward receiver signature
	if err := s.btcClient.CheckRewardAddressSignature(stakeInfo.StakerPublicKey, stakeInfo.ReceiverAddress,
		stakeInfo.ReceiverAddressSignature, stakeInfo.Timestamp,
	); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	log.Info().Msgf("[HTTP] SubmitProofOfStake: %+v ", stakeInfo)
	if stakeInfo.Duration > 0 {
		stakeInfo.Duration = stakeInfo.Duration * 24 * time.Hour.Nanoseconds()
	}

	err := s.backend.CreateStake(stakeInfo.Staker,
		stakeInfo.StakerPublicKey,
		stakeInfo.TxID,
		stakeInfo.Duration,
		stakeInfo.Amount,
		stakeInfo.ReceiverAddress,
		stakeInfo.Reward,
		stakeInfo.ReceiverAddressSignature,
		stakeInfo.Timestamp)
	if err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusInternalServerError, respData)
		return
	}

	c.JSON(http.StatusOK, respData)
}

package api

import (
	"net/http"

	"github.com/generativelabs/btcserver/internal"
	"github.com/gin-gonic/gin"
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

	var staker internal.StakerReq
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

	srakeList := make([]*internal.StakeInfoResp, 0)
	for _, s := range stakes {
		srakeList = append(srakeList, &internal.StakeInfoResp{
			s.Staker,
			s.Tx,
			s.Start,
			s.Duration,
			s.Amount,
			s.RewardReceiver,
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

	var stakeinfo internal.StakeInfoReq
	if err := c.Bind(&stakeinfo); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	//  todo verify ReceiverSignature

	// todo verify btc lock

	// storage
	err := s.backend.CreateStake(stakeinfo.Staker,
		stakeinfo.StakerPublicKey,
		stakeinfo.TxID,
		stakeinfo.Start,
		stakeinfo.Duration,
		stakeinfo.Amount,
		stakeinfo.RewardReceiver,
		stakeinfo.BtcSignature,
		stakeinfo.ReceiverSignature,
		stakeinfo.Timestamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, respData)
}

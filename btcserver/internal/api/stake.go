package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPRespose struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func JSONResp(c *gin.Context, code int, res *HTTPRespose) {
	c.JSON(code, res)
}

type StakerReq struct {
	Staker string `form:"staker" binding:"required"`
	Page   int    `form:"page" binding:"required"`
	Size   int    `form:"size" binding:"required"`
}

type StakeInfoReq struct {
	Staker            string `form:"staker" json:"staker,omitempty"`
	TxID              string `json:"tx_id,omitempty"`
	Start             uint64 `json:"start,omitempty"`
	Duration          uint64 `json:"duration,omitempty"`
	Amount            uint64 `json:"amount,omitempty"`
	RewardReceiver    string `json:"reward_receiver,omitempty"`
	BtcSignature      string `json:"btc_signature,omitempty"`
	ReceiverSignature string `json:"receiver_signature,omitempty"`
	Timestamp         uint64 `json:"timestamp,omitempty"`
}

type StakeInfoResq struct {
	Staker         string `json:"staker,omitempty"`
	Tx             string `json:"tx,omitempty"`
	Start          uint64 `json:"start,omitempty"`
	Durnation      uint64 `json:"durnation,omitempty"`
	Amount         uint64 `json:"amount,omitempty"`
	RewardReceiver string `json:"reward_receiver,omitempty"`
}

// GetStakeListByStaker Get staking history
func (s Server) GetStakeListByStaker(c *gin.Context) {
	respData := &HTTPRespose{Msg: "Ok"}

	var staker StakerReq
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

	srakeList := make([]*StakeInfoResq, 0)
	for _, s := range stakes {
		srakeList = append(srakeList, &StakeInfoResq{
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
func (s Server) SubmitProofOfStake(c *gin.Context) {
	respData := &HTTPRespose{Msg: "Ok"}

	var stakeinfo StakeInfoReq
	if err := c.Bind(&stakeinfo); err != nil {
		respData.Msg = err.Error()
		JSONResp(c, http.StatusBadRequest, respData)
		return
	}

	// todo verify BtcSignature

	// todo verify ReceiverSignature

	// todo verify btc lock

	// storage
	err := s.backend.CreateStake(stakeinfo.Staker,
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

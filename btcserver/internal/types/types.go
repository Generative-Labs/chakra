package types

import "time"

const (
	TimeWheelSize = 5 * time.Minute
)

type StakerReq struct {
	Staker string `form:"staker" binding:"required"`
	Page   int    `form:"page" binding:"required"`
	Size   int    `form:"size" binding:"required"`
}

type StakeInfoReq struct {
	Staker            string `form:"staker" json:"staker,omitempty"`
	StakerPublicKey   string `form:"staker_public_key" json:"staker_public_key,omitempty"`
	TxID              string `json:"tx_id,omitempty"`
	Start             int64  `json:"start,omitempty"`
	Duration          int64  `json:"duration,omitempty"`
	Amount            int64  `json:"amount,omitempty"`
	Reward            int64  `json:"reward,omitempty"`
	RewardReceiver    string `json:"reward_receiver,omitempty"`
	ReceiverSignature string `json:"receiver_signature,omitempty"`
	Timestamp         int64  `json:"timestamp,omitempty"`
}

type StakeInfoResp struct {
	Staker         string `json:"staker,omitempty"`
	Tx             string `json:"tx,omitempty"`
	Start          int64  `json:"start,omitempty"`
	Durnation      int64  `json:"durnation,omitempty"`
	Amount         int64  `json:"amount,omitempty"`
	RewardReceiver string `json:"reward_receiver,omitempty"`
}

type ReleaseTxsInfo struct {
	Staker        string `json:"Staker,omitempty"`
	Tx            string `json:"Tx,omitempty"`
	ReleasingTime int64  `json:"ReleasingTime,omitempty"`
}

type StakeRecordStatus int

const (
	// TxPending defines the StakeRecord status where the Tx has not been included in a block yet.
	TxPending StakeRecordStatus = iota
	// TxIncluded defines the StakeRecord status where the Tx has been included in a block.
	TxIncluded
	// TxFinalized defines the StakeRecord status where the Tx has been confirmed.
	TxFinalized
	// Mismatch defines the StakeRecord status where the Tx in the record does not match the content.
	Mismatch
)

type StakeVerificationParam struct {
	TxID            string
	StakerPubKey    string
	Amount          uint64
	Duration        uint64
	FinalizedStatus StakeRecordStatus
}

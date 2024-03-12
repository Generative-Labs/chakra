package internal

type StakerReq struct {
	Staker string `form:"staker" binding:"required"`
	Page   int    `form:"page" binding:"required"`
	Size   int    `form:"size" binding:"required"`
}

type StakeInfoReq struct {
	Staker            string `form:"staker" json:"staker,omitempty"`
	StakerPublicKey   string `form:"staker_public_key" json:"staker_public_key,omitempty"`
	TxID              string `json:"tx_id,omitempty"`
	Start             uint64 `json:"start,omitempty"`
	Duration          uint64 `json:"duration,omitempty"`
	Amount            uint64 `json:"amount,omitempty"`
	Reward            uint64 `json:"reward,omitempty"`
	RewardReceiver    string `json:"reward_receiver,omitempty"`
	BtcSignature      string `json:"btc_signature,omitempty"`
	ReceiverSignature string `json:"receiver_signature,omitempty"`
	Timestamp         uint64 `json:"timestamp,omitempty"`
}

type StakeInfoResp struct {
	Staker         string `json:"staker,omitempty"`
	Tx             string `json:"tx,omitempty"`
	Start          uint64 `json:"start,omitempty"`
	Durnation      uint64 `json:"durnation,omitempty"`
	Amount         uint64 `json:"amount,omitempty"`
	RewardReceiver string `json:"reward_receiver,omitempty"`
}

type ReleaseTxsInfo struct {
	TxID          string
	ReleasingTime uint64
}

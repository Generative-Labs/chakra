package utils

const (
	unitToken         = 10000 // Unit token 0.0001 BTC corresponding to satoshis
	unitStakingTime   = 24    // Unit pledge time 24 hours
	rewardCoefficient = 0.01  // Reward coefficient
)

// CalculateReward Calculate the reward after staking BTC.
func CalculateReward(totalStakedTokens float64, stakingTimeHours float64) float64 {
	return (totalStakedTokens / unitToken) * (stakingTimeHours / unitStakingTime) * rewardCoefficient
}

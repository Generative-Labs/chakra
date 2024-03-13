package utils

import "time"

// GetUnixEpoch converts a time into a unix timestamp with nanoseconds
func GetUnixEpochFrom(now time.Time) int64 {
	return now.UnixNano()
}

// GetUnixEpoch returns the current time in unix timestamp with the integer part
// representing seconds and the decimal part representing subseconds
func GetUnixEpoch() int64 {
	return GetUnixEpochFrom(time.Now())
}

// Return now time in ms
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Return now time in nano second
func MakeNanoTimestamp() int64 {
	return time.Now().UnixNano()
}

func MakeFloat64Timestamp() float64 {
	return float64(time.Now().UnixNano()) / (float64(time.Millisecond) / float64(time.Nanosecond))
}

func TimeTOTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func TimeToDailyFixedTime(t time.Time) time.Time {
	fixedTime := time.Date(0, 0, 0, t.Hour(), t.Minute(), t.Second(), 0, time.Local)
	return fixedTime
}

func DailyFixedTimeToTime(t time.Time) time.Time {
	now := time.Now()
	fixedTime := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local)

	return fixedTime
}

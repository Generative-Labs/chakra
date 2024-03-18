package api

import (
	"testing"
	"time"
)

func TestInitTimeWheel(t *testing.T) {
	t.Logf("TimeNow %s", time.Now())
	t.Logf("TimeWheel %s", InitTimeWheel())
}

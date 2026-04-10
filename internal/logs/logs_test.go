package logs

import (
	"testing"
)

func TestLogs_New(_ *testing.T) {
	logger := New(false)
	logger.Print("should not see this")
}

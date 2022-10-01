package logs

import (
	"testing"
)

func TestLogs_New(t *testing.T) {
	logger := New(false)
	logger.Print("should not see this")
}

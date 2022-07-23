package logs

import (
	"io"
	"log"
	"os"
)

// New creates a new Logger that only outputs if verbose is set.
//
// Log output is sent to stderr.
func New(verbose bool) *log.Logger {
	if verbose {
		return log.New(os.Stderr, "", log.LstdFlags)
	}
	return log.New(io.Discard, "", 0)
}

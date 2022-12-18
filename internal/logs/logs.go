package logs

import (
	"io"
	"log"
	"os"
)

// New creates a new Logger that only outputs if verbose is set.
//
// Log output is sent to stdout.
func New(verbose bool) *log.Logger {
	if verbose {
		return log.New(os.Stdout, "", log.LstdFlags)
	}
	return log.New(io.Discard, "", 0)
}

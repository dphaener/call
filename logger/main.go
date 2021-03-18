package logger

import (
	"fmt"
	"os"

	"github.com/tj/go-cli-log"
)

var Info = log.Info
var Warn = log.Warn

// Log an error message and exit.
func ExitWithMessage(err error) {
	log.Error(err)
	os.Exit(1)
}

// Print a new line
func NewLine() {
	fmt.Println("")
}

package utils

import (
	"fmt"
	"os"
)

func fatalWithMessage(message string) {
	_, _ = fmt.Fprintf(os.Stderr, message)
	os.Exit(1)
}

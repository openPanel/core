//go:build !linux

package bootstrap

import (
	"log"
)

// Should not compile on non-linux systems.
func init() {
	log.Fatal("This application is only supported on Linux.")
}

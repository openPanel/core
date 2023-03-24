package bootstrap

import (
	"log"
	"os/user"
)

func init() {
	RequireRoot()
}

func RequireRoot() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if u.Uid != "0" {
		log.Fatal("This program requires root privileges")
	}
}

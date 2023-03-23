package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("missing command")
	}

	switch os.Args[1] {
	case "ent":
		ent()
	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}
}

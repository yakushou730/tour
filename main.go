package main

import (
	"log"

	"github.com/yakushou730/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

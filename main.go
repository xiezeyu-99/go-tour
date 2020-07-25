package main

import (
	"github.com/xiezeyu-99/go-programming-tour-book/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

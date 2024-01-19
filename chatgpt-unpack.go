package main

import (
	"fmt"
	"log"
	"os"
)

import "github.com/buger/jsonparser"

func main() {
	text, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
}

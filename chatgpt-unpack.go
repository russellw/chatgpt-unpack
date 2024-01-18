package main

import (
    "os"
    "fmt"
    "log"
)

//import "github.com/buger/jsonparser"

func main() {
    text, err := os.ReadFile("file.txt")
    if err != nil {
        log.Fatal( err)
    }
    fmt.Println(string(text))
}

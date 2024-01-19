package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

import "github.com/buger/jsonparser"

func main() {
	text, err := os.ReadFile("conversations.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = jsonparser.ArrayEach(text, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			log.Fatal(err)
		}
		title, err := jsonparser.GetString(value, "title")
		if err != nil {
			log.Fatal(err)
		}
		updateTime, err := jsonparser.GetFloat(value, "update_time")
		if err != nil {
			log.Fatal(err)
		}
		tm := time.Unix(int64(updateTime), 0)
		fmt.Printf("%s:: %s\n", tm.Format("2006-01-02 15:04:05 MST"), title)
		/*
			// Access the 'mapping' object within the conversation
			mapping, _, _, _ := jsonparser.Get(value, "mapping")
			// Iterate over the mapping objects
			jsonparser.ObjectEach(mapping, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				// Access fields of each mapping object
				mappingID := string(key)
				parent, _ := jsonparser.GetString(value, "parent")

				fmt.Printf("Mapping ID: %s\n", mappingID)
				fmt.Printf("Parent: %s\n", parent)

				return nil
			}, "mapping")
		*/
	})
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

import "github.com/buger/jsonparser"

func main() {
	text, err := os.ReadFile("conversations.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = jsonparser.ArrayEach(text, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		// Access fields of each conversation
		title, _ := jsonparser.GetString(value, "title")
		createTime, _ := jsonparser.GetFloat(value, "create_time")
		updateTime, _ := jsonparser.GetFloat(value, "update_time")

		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Create Time: %.6f\n", createTime)
		fmt.Printf("Update Time: %.6f\n", updateTime)

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
	}, "conversations")
	if err != nil {
		log.Fatal(err)
	}
}

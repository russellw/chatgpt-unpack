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

		// heading
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

		// messages
		mapping, _, _, err := jsonparser.Get(value, "mapping")
		if err != nil {
			log.Fatal(err)
		}
		jsonparser.ObjectEach(mapping, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			mappingID := string(key)

			fmt.Printf("Mapping ID: %s\n", mappingID)

			return nil
		}, "mapping")
	})
	if err != nil {
		log.Fatal(err)
	}
}

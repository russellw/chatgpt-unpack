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
		panic(err)
	}
	_, err = jsonparser.ArrayEach(text, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			panic(err)
		}

		// heading
		title, err := jsonparser.GetString(value, "title")
		if err != nil {
			panic(err)
		}
		updateTime, err := jsonparser.GetFloat(value, "update_time")
		if err != nil {
			panic(err)
		}
		tm := time.Unix(int64(updateTime), 0)
		fmt.Printf("%s:: %s\n", tm.Format("2006-01-02 15:04:05 MST"), title)

		// messages
		mapping, _, _, err := jsonparser.Get(value, "mapping")
		if err != nil {
			panic(err)
		}
		jsonparser.ObjectEach(mapping, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			// message
			message, _, _, err := jsonparser.Get(value, "message")
			if err != nil {
				panic(err)
			}
			if len(message) == 4 {
				return nil
			}

			// author
			author, err := jsonparser.GetString(message, "author", "role")
			if err != nil {
				panic(err)
			}
			switch {
			case author == "user":
				fmt.Print("Q: ")
			case author == "assistant":
				fmt.Print("A: ")
			case author == "tool":
				fmt.Print("A: ")
			case author == "system":
				return nil
			default:
				panic(author)
			}

			// parts
			_, err = jsonparser.ArrayEach(message, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				if err != nil {
					panic(err)
				}
				for i := 0; i < len(value); i++ {
					c := value[i]
					if c == '\\' && i+1 < len(value) {
						i++
						c = value[i]
						if c == 'n' {
							c = '\n'
						}
					}
					fmt.Printf("%c", c)
				}
				fmt.Println()
			}, "content", "parts")
			if err != nil {
				log.Println(string(message))
			}
			if author == "assistant" {
				fmt.Println()
			}
			return nil
		})
		fmt.Println()
	})
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"

	"../../flame"
)

func main() {
	channel, _ := flame.Connect("localhost", 2020)
	fmt.Println("Connected!")
	topicHandler := channel.InitializeTopic("topic")
	topicHandler.Publish("hey")
	fmt.Println("hey")
}

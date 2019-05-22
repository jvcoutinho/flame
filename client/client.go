package main

import (
	"fmt"

	"../../flame"
)

func main() {
	channel, err := flame.Connect("localhost", 2020)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected!")
		_, err = channel.InitializeTopic("topic")
		_, err = channel.AccessTopic("topic2")
		fmt.Println(err)
		end := make(chan int)
		<-end
		//chans, _ := topicHandler.Subscribe()
		fmt.Println("hey")
	}
}

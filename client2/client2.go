package main

import (
	"../../flame"
)

func main() {
	channel, _ := flame.Connect("localhost", 2020, "client4")
	streamTopic, _ := channel.InitializeTopic("Stream")
	for {
		streamTopic.Publish("hey", 2)
		streamTopic.Publish("heylow", 1)
	}
	end := make(chan int)
	<-end
}

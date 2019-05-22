package main

import (
	"../../flame"
)

func main() {
	channel, _ := flame.Connect("localhost", 2020, "client3")
	streamTopic, _ := channel.AccessTopic("Stream")
	streamTopic.Publish([]byte("hey"))
	streamTopic.Subscribe()
	end := make(chan int)
	<-end
}

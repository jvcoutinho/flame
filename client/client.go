package main

import (
	"../../flame"
)

func main() {
	flame.Connect("localhost", 2020, "client1")
	end := make(chan int)
	<-end
}

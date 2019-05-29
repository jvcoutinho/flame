package main

import (
	"fmt"

	"../../flame"
)

func main() {
	channel, _ := flame.Connect("localhost", 2020, "client3")
	_, err := channel.Stream("hello", "C:\\Users\\joao-\\Music\\Synaesthesia Auditiva_320kbps_Masters\\04 Rejuvenescência.mp3")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Streaming done.")
	}
	end := make(chan int)
	<-end
}

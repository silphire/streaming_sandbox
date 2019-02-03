package main

import (
	"fmt"
	"os"

	"github.com/Comcast/gots/packet"
)

func main() {
	var pkt = packet.Packet

	file, err = os.Open("test.mp4")
	if err != null {
		fmt.Printf("Cannot open file -- %s", err.Error())
		os.Exit()
	}

}

package main

import (
	"fmt"
	"os"

	"github.com/Comcast/gots/packet"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Not specified input file\n")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot open file -- %s", err.Error())
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("cannot close file\n")
			os.Exit(1)
		}
	}(file)

	var pkt packet.Packet
	for {
		readLen, err := file.Read(pkt[:])
		if readLen <= 0 {
			break
		}

		if err != nil {
			fmt.Printf("packet read error")
			os.Exit(1)
		}

		pid := packet.Pid(&pkt)

		fmt.Printf("read packet identifier: %d\n", pid)
	}
}

package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
	"strconv"
)

func main() {
	baudRate := uint(14400)
	if len(os.Args) > 1 {
		baudRateStr := os.Args[1]
		if len(baudRateStr) != 0 {
			i, _ := strconv.Atoi(baudRateStr)
			baudRate = uint(i)
		}
	}

	portName := "/dev/ttyACM0"
	if len(os.Args) > 2 {
		portName = os.Args[2]
	}

	// Set up options.
	options := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        baudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	// Reopen the port at 9600 to prevent going back to a special mode
	// next time the port is opened
	port.Close()
	options = serial.OpenOptions{
		PortName:        portName,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err = serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	fmt.Printf("Opened %v at %v\n", portName, baudRate)
}

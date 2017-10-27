package main

import (
	"fmt"
	"log"
	"io"
	"github.com/jacobsa/go-serial/serial"
	"encoding/json"
)

func readSerial(h *Hub) {
	options := serial.OpenOptions{
		PortName: "/dev/ttys006",
		BaudRate: 115200,
		DataBits: 8,
		StopBits: 2,
		InterCharacterTimeout:  100,
		//MinimumReadSize: 100,
	}

	port, err := serial.Open(options)

	defer port.Close()

	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	fmt.Println("Started...")

	for {
		buf := make([]byte, 128)

		n, err := port.Read(buf)

		if err != nil {

			if err != io.EOF {
				fmt.Println("Error reading from serial port: ", err)
			}

		} else {

			buf = buf[:n]
			str := fmt.Sprintf("%s", buf)

			var dat map[string]interface{}

			if err := json.Unmarshal(buf, &dat); err != nil {
				panic(err)
			}

			setItem(dat)

			fmt.Println("[Recived data]: ", str)

			h.broadcast <- buf

		}
	}
}

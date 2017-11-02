package main

import (
	"fmt"
	//"log"
	//"io"
	//"github.com/jacobsa/go-serial/serial"
	//"encoding/json"

	"syscall"
	"os"
	"unsafe"
)

const (
	TCSETS = 0x5402
)

//
//func readSerial(h *Hub, p *string) {
//	options := serial.OpenOptions{
//		PortName:              *p,
//		BaudRate:              115200,
//		DataBits:              8,
//		StopBits:              2,
//		InterCharacterTimeout: 1000,
//		//MinimumReadSize: 100,
//	}
//
//	port, err := serial.Open(options)
//
//	defer port.Close()
//
//	if err != nil {
//		log.Fatalf("serial.Open: %v", err)
//	}
//
//	fmt.Println("Started read from: ", *p)
//
//	for {
//		buf := make([]byte, 128)
//
//		n, err := port.Read(buf)
//
//		if err != nil {
//
//			if err != io.EOF {
//				fmt.Println("Error reading from serial port: ", err)
//			}
//
//		} else {
//
//			buf = buf[:n]
//			str := fmt.Sprintf("%s", buf)
//
//			var dat map[string]interface{}
//
//			if err := json.Unmarshal(buf, &dat); err != nil {
//				panic(err)
//			}
//
//			setItem(dat)
//
//			fmt.Println("[Recived data]: ", str)
//
//			h.broadcast <- buf
//
//		}
//	}
//}

func r(h *Hub, p *string) {
	s, _ := os.OpenFile(*p, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)

	t := syscall.Termios{
		Iflag: syscall.IGNPAR,
		Cflag: syscall.CS8 | syscall.CREAD | syscall.CLOCAL | syscall.B115200,

		Ispeed: syscall.B115200,
		Ospeed: syscall.B115200,
	}

	t.Cc[syscall.VMIN] = 0
	t.Cc[syscall.VTIME] = 100

	// syscall
	syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(s.Fd()),
		uintptr(TCSETS),
		uintptr(unsafe.Pointer(&t)),
		0,
		0,
		0)

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)

		if err != nil { // err will equal io.EOF
			//fmt.Println(err)
		} else {
			buf = buf[:n]
			str := fmt.Sprintf("%s", buf)

			//var dat map[string]interface{}
			//
			//if err := json.Unmarshal(buf, &dat); err != nil {
			//	panic(err)
			//}

			//setItem(dat)

			fmt.Println("[Recived data]: ", str)

			h.broadcast <- buf
		}

	}
}

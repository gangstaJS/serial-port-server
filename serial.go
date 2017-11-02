package main

import (
	"fmt"
	//"log"
	//"github.com/jacobsa/go-serial/serial"
	//"encoding/json"

	"syscall"
	"os"
	"unsafe"

	"bufio"
	"io"
)

const (
	TCSETS = 0x5402
)

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

	reader := bufio.NewReader(s)

	for {
		n, _, err := reader.ReadLine()

		if err == io.EOF {
			// fmt.Println(err)
		} else {
			str := fmt.Sprintf("%s", n)

			//var dat map[string]interface{}
			//
			//if err := json.Unmarshal(n, &dat); err != nil {
			//	panic(err)
			//}
			//
			//setItem(dat)

			fmt.Println("[Recived data]: ", str)

			h.broadcast <- []byte(n)
		}

	}
}

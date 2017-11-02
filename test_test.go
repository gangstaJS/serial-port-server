package main

import (
	"os"
	"syscall"
	"unsafe"
	"fmt"
)

const (
	TCSETS = 0x5402
)

func main() {
	s, _ := os.OpenFile("/dev/ttys006", syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)

	t := syscall.Termios{
		Iflag: syscall.IGNPAR,
		Cflag: syscall.CS8 | syscall.CREAD | syscall.CLOCAL | syscall.B115200,

		//Cc: [20]uint8{syscall.VMIN: 0, syscall.VTIME: 100},

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

	// Send message
	s.Write([]byte("Test message"))

	// Receive reply

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)

		if err != nil { // err will equal io.EOF
			fmt.Println(err)
		}

		fmt.Printf("%v\n", string(buf[:n]))
	}

}

package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return
	}
	var addr = syscall.SockaddrInet4{
		Port: 1234,
		Addr: [4]byte{127, 0, 0, 1},
	}
	syscall.Connect(fd, &addr)

	msg := []byte("Hello")
	syscall.Write(fd, msg)

	p := make([]byte, 4096)
	read, err := syscall.Read(fd, p)
	if err != nil {
		return
	}
	if read < 0 {
		fmt.Printf("read error %v", err)
		return
	}
	fmt.Printf("Server says %s\n", string(p[:]))
}

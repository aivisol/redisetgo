package main

import (
	"fmt"
	"syscall"
)

type sockaddr_in struct {
	sinFamily int
	sinPort   int
	sinAddr   int
}

func main() {

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Printf("socket()")
		return
	}
	val := 1
	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, val)
	if err != nil {
		fmt.Printf("error SetsockoptByte(): %v\n", err)
		return
	}
	var addr = syscall.SockaddrInet4{
		Port: 1234,
		Addr: [4]byte{0, 0, 0, 0},
	}

	err = syscall.Bind(fd, &addr)
	if err != nil {
		return
	}

	err = syscall.Listen(fd, syscall.SOMAXCONN)
	if err != nil {
		return
	}

	for {
		nfd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Printf("Accept error %v\n", err)
			return
		}
		doSomething(nfd)
		syscall.Close(fd)
	}
}

func doSomething(nfd int) {
	p := make([]byte, 4096)
	read, err := syscall.Read(nfd, p)
	if err != nil {
		fmt.Printf("read error %v", err)
		return
	}
	if read < 0 {
		fmt.Printf("read error")
		return
	}
	fmt.Printf("%d bytes read\n", read)

	fmt.Printf("Client says %s\n", string(p[:]))
	wbuf := []byte("world")
	n, err := syscall.Write(nfd, wbuf)
	if err != nil {
		fmt.Printf("write error %v", err)
		return
	}
	fmt.Printf("%d bytes written\n", n)
}

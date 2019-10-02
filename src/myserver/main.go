package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Logged: ", ln)
		//n, err := fmt.Fprintf(conn, " Scanned: ", ln)
		/*if err != nil {
			log.Println(err)
		}
		fmt.Println(n)*/
	}
	defer conn.Close()
	fmt.Println("Connection closed...")
}

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
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()
		go request(conn)
		go response(conn)
	}

}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><head></head><body>Hello World from GoLang</body><html>`
	_, _ = fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	_, _ = fmt.Fprintf(conn, "Content-Length: %d \r\n", len(body))
	_, _ = fmt.Fprintf(conn, "Content-Type: text/html \r\n")
	_, _ = fmt.Fprintf(conn, "\r\n")
	_, _ = fmt.Fprintf(conn, body)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
}

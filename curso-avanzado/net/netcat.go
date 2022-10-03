package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 3090, "Port to open")
	host = flag.String("host", "localhost", "Host listening")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContent(dest io.Writer, src io.Reader) {
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Message struct {
	Body string
}

var listenAddr = flag.String("listen", "localhost:8002", "host:port to listen on")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}
		go serve(conn, os.Stdout)
	}
}

func GetMessage(s string) Message {
	return Message{s}
}

func serve(conn io.ReadCloser, w io.Writer) {
	defer conn.Close()

	dec := json.NewDecoder(conn)
	for {
		message := GetMessage("")
		err := dec.Decode(&message)
		if err != nil {
			//log.Fatal(err)
			log.Println(err)
			return
		}
		fmt.Fprint(w, message.Body)
	}
}

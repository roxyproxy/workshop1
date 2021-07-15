package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "localhost:8002", "host:port to dial")

type Message struct {
	Body string
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *dialAddr)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		message := GetMessage(scanner.Text())
		//err := SendMessage(message, os.Stdout)
		err := SendMessage(message, conn)
		if err != nil {
			log.Fatal(err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	conn.Close()
}

func GetMessage(s string) Message {
	return Message{s}
}

func SendMessage(m Message, writer io.Writer) error {
	enc := json.NewEncoder(writer)
	return enc.Encode(m)
}

// Solution to part 6 of the Whispering Gophers code lab.
//
// This program is functionally equivalent to part 5,
// but the reading from standard input and writing to the
// network connection are done by separate goroutines.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/campoy/whispering-gophers/util"
)

var (
	peerAddr = flag.String("peer", "", "peer host:port")
	self     string
)

type Message struct {
	Addr string
	Body string
}

func main() {
	flag.Parse()

	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	self = l.Addr().String()
	log.Println("Listening on", self)

	// TODO: Make a new channel of Messages.
	var messages = make(chan Message)

	go dial(*peerAddr, messages)
	go readInput(messages, os.Stdin, self)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	d := json.NewDecoder(c)
	for {
		var m Message
		err := d.Decode(&m)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%#v\n", m)
	}
}

func readInput(ch chan<- Message, r io.Reader, selfAddr string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		m := Message{
			Addr: selfAddr,
			Body: s.Text(),
		}
		// TODO: Send the message to the channel of messages.
		ch <- m
	}
	//close(messages)
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func dial(addr string, ch <-chan Message) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(addr, err)
		return
	}
	defer c.Close()

	e := json.NewEncoder(c)

	for m := range ch {
		err := e.Encode(m)
		if err != nil {
			log.Println(addr, err)
			return
		}
	}
}

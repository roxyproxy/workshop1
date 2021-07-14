package main

import (
	"encoding/json"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReadInput(t *testing.T) {
	want := Message{Addr: "192.168.1.8:58419", Body: "Test Body"}
	m := make(chan Message, 1)
	done := make(chan interface{})

	go func() {
		defer func() { close(done) }()
		r := strings.NewReader("Test Body")
		readInput(m, r, "192.168.1.8:58419")
	}()

	<-done
	select {
	case got := <-m:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	default:
		t.Errorf(`didn't receive message on channel`)
	}
}

func TestDial(t *testing.T) {
	m := make(chan Message, 1)
	want := Message{Addr: "192.168.1.8:58419", Body: "Test Body"}
	m <- want
	close(m)
	done := make(chan bool)
	addr := "localhost:8006"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	go func() {
		defer func() { done <- true }()

		conn, err := l.Accept()
		defer conn.Close()

		if err != nil {
			t.Error(err)
		}

		d := json.NewDecoder(conn)
		for {
			var got Message
			err := d.Decode(&got)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
			return
		}
	}()

	go dial(addr, m)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Errorf("expired timeout")
	}

}

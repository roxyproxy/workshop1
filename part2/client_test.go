package main

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

var MessageBody = "Awesome body!"

func TestGetMessage(t *testing.T) {
	got := GetMessage(MessageBody)
	want := Message{MessageBody}

	assertMessage(t, got, want)
}

func TestSendMessage(t *testing.T) {
	want := `{"Body":"Awesome body!"}`
	buf := bytes.Buffer{}
	err := SendMessage(GetMessage(MessageBody), &buf)
	if err != nil {
		t.Errorf("error is not nil %+v", err)
	}
	got := strings.Trim(buf.String(), "\n")
	assertMessage(t, got, want)
}

func assertMessage(t *testing.T, got interface{}, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func ExampleMessage() {
	message := GetMessage(MessageBody)
	SendMessage(message, os.Stdout)
	// Output: {"Body":"Awesome body!"}
}

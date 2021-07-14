package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestProceedMessage(t *testing.T) {

	want := `{"Addr":"192.168.1.8:58419","Body":"Test Body"}`
	r := strings.NewReader("Test Body")
	buf := bytes.Buffer{}
	err := proceedMessage(r, &buf, "192.168.1.8:58419")
	if err != nil {
		t.Errorf("error is not nil %+v", err)
	}
	got := strings.Trim(buf.String(), "\n")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var MessageBody = "Awesome body!"

type stubReadCloser struct {
	closed bool
}

func (s *stubReadCloser) Close() error {
	s.closed = true
	return nil
}

func (s stubReadCloser) Read(b []byte) (n int, err error) {

	return 0, io.EOF
}

func TestServe(t *testing.T) {
	t.Run("returns an error if got != want", func(t *testing.T) {
		want := "Awesome body!"
		m := GetMessage(MessageBody)
		b, _ := json.Marshal(&m)
		r := ioutil.NopCloser(bytes.NewBuffer(b))
		buf := bytes.Buffer{}
		serve(r, &buf)
		got := buf.String()
		assertMessage(t, got, want)
	})

	t.Run("returns an error id connection is not closed", func(t *testing.T) {
		want := true
		r := stubReadCloser{}
		serve(&r, os.Stdout)
		//.Println(r.Closed)
		got := r.closed
		assertMessage(t, got, want)
	})

}

func assertMessage(t *testing.T, got interface{}, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

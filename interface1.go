package main

import (
	"fmt"
	"bytes"
	"io "
)

func main() {
	var myOBj interface{} = NewBufferedWriterCloser()
	// var w Writer = ConsoleWriter{}
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("hello this is a test"))
	wc.Close()

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed")
	}

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

type (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, error := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err := fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nill
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*
Interfaces are named collections of method signatures.
*/

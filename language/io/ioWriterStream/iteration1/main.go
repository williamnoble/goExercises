package main

import (
	"fmt"
	"io"
	"os"
)

type channelWriter struct {
	Channel chan byte
}

func NewChannelWriter() *channelWriter {
	return &channelWriter{
		Channel: make(chan byte, 1024),
	}
}

func (c *channelWriter) Write(p []byte) (int, error) {
	// wrote the contents of p, now p == 0
	if len(p) == 0 {
		return 0, nil
	}

	go func() {
		defer close(c.Channel)
		for _, b := range p {
			c.Channel <- b
		}
	}()
	return len(p), nil
}

func main() {
	// Create a 1024 []byte chan
	chanWriter := NewChannelWriter()

	// Write to that chan
	go func() {
		fmt.Fprint(chanWriter, "Stream me")
	}()

	// Block. Read from the chan
	for c := range chanWriter.Channel {
		fmt.Printf("%c \n", c)
	}

	chanWriterF := NewChannelWriter()
	file, err := os.Open("./file.dat")
	if err != nil {
		fmt.Println("Error reading the file", err)
		os.Exit(1)
	}
	// io.Copy (dst, src) - Copy from file -> chanWriterF
	_, err = io.Copy(chanWriterF, file)
	if err != nil {
		fmt.Println("Error copying", err)
		os.Exit(1)
	}
	for c := range chanWriterF.Channel {
		fmt.Printf("%c\n", c)
	}
}

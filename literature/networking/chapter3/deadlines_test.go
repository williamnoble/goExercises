package ch03

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	//
	sync := make(chan struct{})

	// setup tcp server
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer func() {
			//goland:noinspection GoUnhandledErrorResult
			conn.Close()
			close(sync)
		}()

		// Once client connected to server set a deadline.
		//
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.Error(err)
			return
		}

		// Create buffer for Conn to read. Read blocks until it reads data from node.
		// Read block lifted once deadline is passed.
		buf := make([]byte, 1)
		_, err = conn.Read(buf)     // blocks and fails as conn does not send data
		nErr, ok := err.(net.Error) // Check for timeout (inevitable -> as no data sent)
		if !ok || !nErr.Timeout() {
			t.Errorf("expected timeout error; actual %v", err)
		}

		sync <- struct{}{}

		//  Reset the deadline!
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.Error(err)
			return
		}

		_, err = conn.Read(buf)
		if err != nil {
			t.Error(err)
		}

	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	<-sync

	_, err = conn.Write([]byte("1"))
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1)
	_, err = conn.Read(buf)
	if err != io.EOF {
		t.Errorf("expected server termination; actual: %v", err)
	}
}

package ch03

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:2400")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("bound to %q", listener.Addr())
	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()

		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			go func(c net.Conn) {
				defer func() {
					//goland:noinspection GoUnhandledErrorResult
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal("error here :", err)
	}

	//goland:noinspection GoUnhandledErrorResult
	conn.Close()
	<-done
	//goland:noinspection GoUnhandledErrorResult
	listener.Close()
	<-done

}

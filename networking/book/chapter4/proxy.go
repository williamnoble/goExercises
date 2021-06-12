package ch04

import (
	"io"
	"net"
	"sync"
	"testing"
)

func proxy(from io.Reader, to io.Writer) error {
	// this is lie asking fromWriter, ok := from.(io.Writer)
	// i.e. does fromWriter conform to io.Writer. fromWriter then
	// holds the underlying value
	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if toIsReader && fromIsWriter {
		go func() {
			_, _ = io.Copy(fromWriter, toReader)
		}()
	}

	return nil
}

func TestProxy(t *testing.T) {
	var wg sync.WaitGroup
	// server listens for 'ping' and replies with 'pong'

	server, err := net.Listen("tcp", "127.0.0.1":)
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				return
			}

			go func(c net.Conn) {
				defer c.Close()
				for {
					buf := make([]byte, 1024)
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}

					switch msg := string(buf[:n]); msg {
					case "ping":
						_, err = c.Write([]byte("pong"))
					default:
						_, err = c.Write(buf[:n])
					}
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
				}
			}(conn)
		}

	}()

	proxyServer, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			conn, err := proxyServer.Accept()
		}
	}()
}

package ch04

import (
	"io"
	"net"
)

func proxyConn(source, destination string) error {
	connSource, err := net.Dial("tpc", source)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		defer connSource.Close()
	}

	connDestination, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer connDestination.Close()

	// connDestination replies to connSource
	go func() {
		_, _ = io.Copy(connSource, connDestination)
	}()

	return err
}

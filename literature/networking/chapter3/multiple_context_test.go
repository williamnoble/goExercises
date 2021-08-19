package ch03

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"
)

func TestDialContextCancelFanout(t *testing.T) {
	// Setup Context with 10s deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	// Setup tcp server
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer listener.Close()

	// Accept one client only
	go func() {
		conn, err := listener.Accept()
		if err == nil {
			//goland:noinspection GoUnhandledErrorResult
			conn.Close()
		}
	}()

	// Setup client
	dial := func(ctx context.Context, address string, response chan int, id int, wg *sync.WaitGroup) {
		defer wg.Done()
		var d net.Dialer

		c, err := d.DialContext(ctx, "tcp", address)
		if err != nil {
			t.Logf("client with id:%d failed to connect", id)
			return

		}
		//goland:noinspection GoUnhandledErrorResult
		c.Close()

		// DialContext will fail if the first (random) Client has connected
		// Therefore flow blocks here:
		// Case1 : ctx.Done = ctx Channel is closed
		// Case2: receive the id of client which failed to connect
		select {
		case <-ctx.Done():
		case response <- id:
		}
	}

	resp := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dial(ctx, listener.Addr().String(), resp, i+1, &wg)
	}
	responseID := <-resp
	cancel()
	wg.Wait()
	close(resp)

	if ctx.Err() != context.Canceled {
		t.Errorf("expected cancelled context; actual: %v", ctx.Err())
	}
	t.Logf("dialer %d retrieved the resource", responseID)
}

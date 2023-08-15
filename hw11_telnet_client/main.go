package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type contextKey string

var (
	timeout       time.Duration
	errContextKey contextKey = "err"
)

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "set connection timeout")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		os.Stderr.WriteString("expected at least 2 arguments")
		os.Exit(1)
	}

	client := NewTelnetClient(net.JoinHostPort(args[0], args[1]), timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("cannot connect to server %s", err))
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(2)

	go receive(&ctx, cancel, client, &wg)
	go send(&ctx, cancel, client, &wg)

	wg.Wait()
}

func receive(ctx *context.Context, cancel context.CancelFunc, client TelnetClient, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-(*ctx).Done():
			return
		default:
			if err := client.Receive(); err != nil {
				closeConnection(ctx, cancel, client, err)
			}
		}
	}
}

func send(ctx *context.Context, cancel context.CancelFunc, client TelnetClient, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-(*ctx).Done():
			return
		default:
			if err := client.Send(); err != nil {
				closeConnection(ctx, cancel, client, err)
			}
		}
	}
}

func closeConnection(ctx *context.Context, cancel context.CancelFunc, client TelnetClient, err error) {
	client.Close()
	*ctx = context.WithValue(*ctx, errContextKey, err)
	cancel()
}

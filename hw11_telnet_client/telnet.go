package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &TCPClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

type TCPClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

var _ TelnetClient = (*TCPClient)(nil)

func (t *TCPClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}

func (t *TCPClient) Close() error {
	return t.conn.Close()
}

func (t *TCPClient) Send() error {
	_, err := io.Copy(t.conn, t.in)
	return err
}

func (t *TCPClient) Receive() error {
	_, err := io.Copy(t.out, t.conn)
	return err
}

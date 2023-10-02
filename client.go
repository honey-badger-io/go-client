package hb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	sys  SysClient
	data DataClient
}

func (c *Client) Data() DataClient {
	return c.data
}

func (c *Client) Ping(ctx context.Context) (string, error) {
	res, err := c.sys.Ping(ctx, &PingRequest{})
	if err != nil {
		return "", err
	}

	return res.Mesage, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func NewClient(target string) (*Client, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &Client{
		sys:  NewSysClient(conn),
		data: NewDataClient(conn),
		conn: conn,
	}

	return client, nil
}

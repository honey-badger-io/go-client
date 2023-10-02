package hb

import (
	"context"

	"github.com/honey-badger-io/go-client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	sys  pb.SysClient
	data *Data
}

func (c *Client) Data() *Data {
	return c.data
}

func (c *Client) Ping(ctx context.Context) (string, error) {
	res, err := c.sys.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		return "", err
	}

	return res.Mesage, nil
}

func NewClient(target string) (*Client, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	data := &Data{
		grpc: pb.NewDataClient(conn),
	}

	client := &Client{
		sys:  pb.NewSysClient(conn),
		data: data,
		conn: conn,
	}

	return client, nil
}

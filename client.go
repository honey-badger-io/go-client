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
	data pb.DataClient
	db   pb.DbClient
}

func (c *Client) Data(ctx context.Context, db string) *Data {
	return &Data{
		c:   c.data,
		ctx: ctx,
		db:  db,
	}
}

func (c *Client) Db(ctx context.Context, db string) *DbClient {
	return &DbClient{
		d:   c.db,
		ctx: ctx,
		db:  db,
	}
}

func (c *Client) Ping(ctx context.Context) (string, error) {
	res, err := c.sys.Ping(ctx, &pb.PingRequest{})
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
		sys:  pb.NewSysClient(conn),
		data: pb.NewDataClient(conn),
		db:   pb.NewDbClient(conn),
		conn: conn,
	}

	return client, nil
}

package hb

import (
	"context"

	"github.com/honey-badger-io/go-client/pb"
)

type Data struct {
	grpc pb.DataClient
}

func (d *Data) Get(ctx context.Context, db string, key string) (*pb.GetResult, error) {
	return d.grpc.Get(ctx, &pb.KeyRequest{
		Db:  db,
		Key: key,
	})
}

func (d *Data) Set(ctx context.Context, req *pb.SetRequest) error {
	_, err := d.grpc.Set(ctx, req)

	return err
}

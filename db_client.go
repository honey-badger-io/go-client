package hb

import (
	"context"

	"github.com/honey-badger-io/go-client/pb"
)

type DbClient struct {
	ctx context.Context
	d   pb.DbClient
	db  string
}

type CreateDbOptions struct {
	InMemory bool
}

func (d *DbClient) Create(opt CreateDbOptions) error {
	_, err := d.d.Create(d.ctx, &pb.CreateDbReq{
		Name: d.db,
		Opt: &pb.CreateDbOpt{
			InMemory: opt.InMemory,
		},
	})

	return err
}

func (d *DbClient) Drop() error {
	_, err := d.d.Drop(d.ctx, &pb.DropDbRequest{
		Name: d.db,
	})

	return err
}

func (d *DbClient) Exists() (bool, error) {
	res, err := d.d.Exists(d.ctx, &pb.ExistsDbReq{
		Name: d.db,
	})

	return res.Exists, err
}

func (d *DbClient) EnsureDb(opt CreateDbOptions) error {
	_, err := d.d.EnsureDb(d.ctx, &pb.CreateDbReq{
		Name: d.db,
		Opt: &pb.CreateDbOpt{
			InMemory: opt.InMemory,
		},
	})

	return err
}

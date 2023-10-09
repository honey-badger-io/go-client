package hb

import (
	"context"

	"github.com/honey-badger-io/go-client/pb"
)

type Data struct {
	ctx context.Context
	c   pb.DataClient
	db  string
}

func (d *Data) Get(key string) ([]byte, bool, error) {
	res, err := d.c.Get(d.ctx, &pb.KeyRequest{
		Db:  d.db,
		Key: key,
	})

	if err != nil {
		return nil, false, err
	}

	return res.Data, res.Hit, nil
}

func (d *Data) Set(key string, data []byte) error {
	_, err := d.c.Set(d.ctx, &pb.SetRequest{
		Db:   d.db,
		Key:  key,
		Data: data,
	})
	return err
}

func (d *Data) SetWithTtl(key string, data []byte, ttl int32) error {
	_, err := d.c.Set(d.ctx, &pb.SetRequest{
		Db:   d.db,
		Key:  key,
		Data: data,
		Ttl:  &ttl,
	})
	return err
}

func (d *Data) Delete(key string) error {
	_, err := d.c.Delete(d.ctx, &pb.KeyRequest{
		Db:  d.db,
		Key: key,
	})
	return err
}

func (d *Data) DeleteByPrefix(prefix string) error {
	_, err := d.c.DeleteByPrefix(d.ctx, &pb.PrefixRequest{
		Db:     d.db,
		Prefix: prefix,
	})
	return err
}

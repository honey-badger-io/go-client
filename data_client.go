package hb

import (
	"context"
	"io"

	"github.com/honey-badger-io/go-client/pb"
)

type Data struct {
	ctx context.Context
	c   pb.DataClient
	db  string
}

type ReadCallback func(key string, data []byte) error

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

func (d *Data) NewSendStream() (*SendStream, error) {
	stream, err := d.c.CreateSendStream(d.ctx)
	if err != nil {
		return nil, err
	}

	// First message need to have database name
	err = stream.Send(&pb.SendStreamReq{
		Db: d.db,
	})
	if err != nil {
		return nil, err
	}

	return &SendStream{s: stream}, nil
}

func (d *Data) Read(prefix string, callback ReadCallback) error {
	stream, err := d.c.CreateReadStream(d.ctx, &pb.ReadStreamReq{
		Db:     d.db,
		Prefix: &prefix,
	})
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	for {
		itm, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := callback(itm.Key, itm.Data); err != nil {
			return err
		}
	}
}

package hb

import "context"

type Data struct {
	ctx context.Context
	c   DataClient
	db  string
}

func (d *Data) Get(key string) (*GetResult, error) {
	return d.c.Get(d.ctx, &KeyRequest{
		Db:  d.db,
		Key: key,
	})
}

func (d *Data) Set(key string, data []byte) error {
	_, err := d.c.Set(d.ctx, &SetRequest{
		Db:   d.db,
		Key:  key,
		Data: data,
	})
	return err
}

func (d *Data) SetWithTtl(key string, data []byte, ttl int32) error {
	_, err := d.c.Set(d.ctx, &SetRequest{
		Db:   d.db,
		Key:  key,
		Data: data,
		Ttl:  &ttl,
	})
	return err
}

func (d *Data) Delete(key string) error {
	_, err := d.c.Delete(d.ctx, &KeyRequest{
		Db:  d.db,
		Key: key,
	})
	return err
}

func (d *Data) DeleteByPrefix(prefix string) error {
	_, err := d.c.DeleteByPrefix(d.ctx, &PrefixRequest{
		Db:     d.db,
		Prefix: prefix,
	})
	return err
}

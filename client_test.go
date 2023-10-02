package hb

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Db = "go-client-tests"
)

func TestPing(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	t.Run("should call Ping", func(t *testing.T) {
		msg, err := client.Ping(context.TODO())

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.Equal(t, "pong", msg)
	})
}

func TestSet(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	t.Run("should set bytes", func(t *testing.T) {
		_, err := client.Data().Set(context.TODO(), &SetRequest{
			Db:   Db,
			Key:  "set-bytes",
			Data: make([]byte, 1),
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})
}

func TestGet(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	t.Run("should get bytes", func(t *testing.T) {
		const key = "get-bytes"
		data := make([]byte, 1)

		client.Data().Set(context.TODO(), &SetRequest{
			Db:   Db,
			Key:  "get-bytes",
			Data: data,
		})

		res, err := client.Data().Get(context.TODO(), &KeyRequest{
			Db:  Db,
			Key: key,
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.Equal(t, data, res.Data)
	})
}

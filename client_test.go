package hb

import (
	"context"
	"fmt"
	"testing"
	"time"

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

	t.Run("should set data", func(t *testing.T) {
		err := client.
			Data(context.TODO(), Db).
			Set("set-bytes", make([]byte, 1))

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})

	t.Run("should set data with ttl", func(t *testing.T) {
		key := "set-bytes"

		err := client.
			Data(context.TODO(), Db).
			SetWithTtl(key, make([]byte, 1), 1)
		time.Sleep(1500 * time.Millisecond)

		res, _ := client.Data(context.TODO(), Db).Get(key)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.False(t, res.Hit)
	})
}

func TestGet(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	t.Run("should get data", func(t *testing.T) {
		const key = "get-bytes"
		value := make([]byte, 1)
		data := client.Data(context.TODO(), Db)

		data.Set(key, value)

		res, err := data.Get(key)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.Equal(t, data, res.Data)
	})
}

func TestDelete(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	t.Run("should delete by key", func(t *testing.T) {
		key := "delete-by-key"
		data := client.Data(context.TODO(), Db)
		data.Set(key, make([]byte, 1))

		err := data.Delete(key)
		res, _ := data.Get(key)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.False(t, res.Hit)
	})

	t.Run("should delete by prefix", func(t *testing.T) {
		prefix := "delete-by-prefix-"
		data := client.Data(context.TODO(), Db)

		data.Set(fmt.Sprintf("%s%d", prefix, 1), make([]byte, 1))
		data.Set(fmt.Sprintf("%s%d", prefix, 2), make([]byte, 2))

		err := data.DeleteByPrefix(prefix)
		res1, _ := data.Get(fmt.Sprintf("%s%d", prefix, 1))
		res2, _ := data.Get(fmt.Sprintf("%s%d", prefix, 2))

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.False(t, res1.Hit)
		assert.False(t, res2.Hit)
	})
}

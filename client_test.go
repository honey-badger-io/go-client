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

var CreateOpt = CreateDbOptions{
	InMemory: true,
}

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

	client.Db(context.TODO(), Db).EnsureDb(CreateOpt)

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

		_, hit, _ := client.Data(context.TODO(), Db).Get(key)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.False(t, hit)
	})
}

func TestGet(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	client.Db(context.TODO(), Db).EnsureDb(CreateOpt)

	t.Run("should get data", func(t *testing.T) {
		const key = "get-bytes"
		value := make([]byte, 1)
		data := client.Data(context.TODO(), Db)

		data.Set(key, value)

		resultData, _, err := data.Get(key)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.Equal(t, value, resultData)
	})
}

func TestDelete(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	client.Db(context.TODO(), Db).EnsureDb(CreateOpt)

	t.Run("should delete by key", func(t *testing.T) {
		key := "delete-by-key"
		data := client.Data(context.TODO(), Db)
		data.Set(key, make([]byte, 1))

		err := data.Delete(key)
		_, hit, _ := data.Get(key)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.False(t, hit)
	})

	t.Run("should delete by prefix", func(t *testing.T) {
		prefix := "delete-by-prefix-"
		data := client.Data(context.TODO(), Db)

		data.Set(fmt.Sprintf("%s%d", prefix, 1), make([]byte, 1))
		data.Set(fmt.Sprintf("%s%d", prefix, 2), make([]byte, 2))

		err := data.DeleteByPrefix(prefix)
		_, hit1, _ := data.Get(fmt.Sprintf("%s%d", prefix, 1))
		_, hit2, _ := data.Get(fmt.Sprintf("%s%d", prefix, 2))

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.False(t, hit1)
		assert.False(t, hit2)
	})
}

func TestSendStream(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	client.Db(context.TODO(), Db).EnsureDb(CreateOpt)

	t.Run("should send data using stream", func(t *testing.T) {
		stream, err := client.Data(context.TODO(), Db).NewSendStream()
		if err != nil {
			panic(err)
		}

		err1 := stream.Send("send-stream-1", make([]byte, 1))
		err2 := stream.Send("send-stream-2", make([]byte, 1))
		errClose := stream.Close()

		assert.Nil(t, err1, fmt.Sprintf("%v", err1))
		assert.Nil(t, err2, fmt.Sprintf("%v", err2))
		assert.Nil(t, errClose, fmt.Sprintf("%v", errClose))
	})
}

func TestReadByPrefix(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	client.Db(context.TODO(), Db).EnsureDb(CreateOpt)

	t.Run("should read data by stream", func(t *testing.T) {
		data := client.Data(context.TODO(), Db)

		stream, _ := data.NewSendStream()
		stream.Send("read-stream-1", make([]byte, 1))
		stream.Send("read-stream-2", make([]byte, 2))
		stream.Close()

		result := make(map[string][]byte)

		err := data.Read("read-stream", func(key string, data []byte) error {
			result[key] = data
			return nil
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.Equal(t, make([]byte, 1), result["read-stream-1"])
		assert.Equal(t, make([]byte, 2), result["read-stream-2"])
	})
}

func TestDb(t *testing.T) {
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}
	const TestDb = "test-create-db"

	t.Run("should call create db", func(t *testing.T) {
		dbClient := client.Db(context.TODO(), TestDb)

		err := dbClient.Create(CreateDbOptions{
			InMemory: true,
		})

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})

	t.Run("should call drop db", func(t *testing.T) {
		dbClient := client.Db(context.TODO(), TestDb)

		err := dbClient.Drop()

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})

	t.Run("should call exists db", func(t *testing.T) {
		dbClient := client.Db(context.TODO(), "test-db-exists")

		dbClient.Create(CreateOpt)
		defer dbClient.Drop()

		result, err := dbClient.Exists()

		assert.Nil(t, err, fmt.Sprintf("%v", err))
		assert.True(t, result)
	})

	t.Run("should call ensure db", func(t *testing.T) {
		dbClient := client.Db(context.TODO(), "test-db-ensure")

		dbClient.Create(CreateOpt)
		defer dbClient.Drop()

		err := dbClient.EnsureDb(CreateOpt)

		assert.Nil(t, err, fmt.Sprintf("%v", err))
	})
}

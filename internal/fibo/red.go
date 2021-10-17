package fibo

import (
	"context"
	"encoding/base32"
	"encoding/binary"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ExampleClient() {
	ctx := context.TODO()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
func keyAsString(key uint64) string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(key))

	return base32.HexEncoding.EncodeToString(b)
}

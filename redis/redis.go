package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)
func main() {
	ctx := context.Background()

    rdb := redis.NewClient(&redis.Options{
    	Addr: "192.168.33.13:6379",
    	Network: "tcp",
    	Password: "myredis",
    },
    )
    rdb.Set(ctx, "hello", "world",0)
	val, err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	rdb.ZAdd(ctx, "fun", &redis.Z{
		Score  : 42,
		Member : "a" ,
	},  &redis.Z{
		Score  : 24,
		Member : "b" ,
	}, &redis.Z{
		Score  : 36,
		Member : "c" ,
	})
	vals, err := rdb.ZRange(ctx,"fun", 0, -1).Result()
	if err != nil{
		panic(err)
	}
	fmt.Printf("%#v\n", vals)

	valz, _ := rdb.ZRangeWithScores(ctx, "fun", 0, -1).Result()
	for _,z := range valz {
		fmt.Println(z.Member, z.Score)
	}

	vale, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()
	fmt.Println(vale)

	pipe := rdb.Pipeline()

	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	// Execute
	//
	//     INCR pipeline_counter
	//     EXPIRE pipeline_counts 3600
	//
	// using one rdb-server roundtrip.
	_, err = pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)
}
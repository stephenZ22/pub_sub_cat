package main

import (
	"context"
	"fmt"
	"pub_sub_cat/util"
)

func main() {
	rdb := util.CreateRedisConn()
	ctx := context.Background()

	pubSub := rdb.Subscribe(ctx, "test-channel")
	for obj := range pubSub.Channel() {
		fmt.Println("aaa")
		fmt.Println(obj.Payload)
		fmt.Printf("currrnt rdb object point addr %p\n", &rdb)

	}
	// time.Sleep(time.Second * 5)
}

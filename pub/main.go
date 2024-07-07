package main

import (
	"context"
	"encoding/json"
	"fmt"
	"pub_sub_cat/model"
	"pub_sub_cat/util"
	"time"
)

func main() {
	rdb := util.CreateRedisConn()
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		fmt.Printf("currrnt rdb object point addr %p\n", &rdb)
		user := model.User{
			Name:  fmt.Sprintf("wangxiao-%d", i),
			Email: fmt.Sprintf("wangxiao-%d@test.com", i),
		}
		userRaw, _ := json.Marshal(&user)
		err := rdb.Publish(ctx, "test-channel", userRaw).Err()
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(time.Second * 1)
	}
}

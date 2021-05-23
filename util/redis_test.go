package yizuutil

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	rdb := GetRedis()
	ctx := rdb.Context()
	type Abc struct {
		A string
		B int
	}
	abc := Abc{}
	abc.A = "hello"
	abc.B = 20
	// 设置的时候要先用JSON处理一下
	b, _ := json.Marshal(abc)
	k, v := rdb.SetEX(ctx, "abc", b, 10 * time.Second).Result()
	if v != nil {
		fmt.Println("hello")
	} else {
		fmt.Println("okkkkkk")
		fmt.Println(k)
	}
	var s Abc
	// 获取不存在的键返回nil
	x, _ := rdb.Get(ctx, "xxx").Bytes()

	// 解析
	json.Unmarshal(x, &s)

	fmt.Println(s)
}
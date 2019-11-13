package test

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

func Redis() cache.Cache {
	redis, error := cache.NewCache("redis", `{"key":"user","conn":"*:6380","dbNum":"0","password":"Odd9qD5TYfZsWg5z"}`)
	if error != nil {
		fmt.Println("redis error:", error)
	}
	return redis
}

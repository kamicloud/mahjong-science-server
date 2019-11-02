package utils

import c "github.com/astaxie/beego/cache"

var Cache c.Cache

func init() {
	bm, _ := c.NewCache("memory", `{"interval":999999}`)
	Cache = bm
}
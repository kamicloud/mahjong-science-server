package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func init() {
	bm := cache.New(5*time.Minute, 10*time.Minute)
	Cache = bm
}
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rpcx"
)

// Config 信息
type Config struct {
	rpcx.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
}

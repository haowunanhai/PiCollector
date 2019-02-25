package rds

import (
	"github.com/garyburd/redigo/redis"
)

// Rdser redis操作接口
type Rdser interface {
	DoCmd(cmd string, args ...interface{}) (interface{}, error)
	DoScript(script *redis.Script, args ...interface{}) (interface{}, error)
}

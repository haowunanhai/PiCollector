package codis

import (
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
)

// Codis 结构
type Codis []*redis.Pool

// New 创建访问Codis的连接池
func New(config Config) Codis {
	c := make(Codis, len(config.Addrs))
	for idx, addr := range config.Addrs {
		c[idx] = &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			IdleTimeout: config.IdleTimeout * time.Millisecond,
			Wait:        config.Wait,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.DialTimeout(
					"tcp",
					addr,
					config.ConnectTimeout*time.Millisecond,
					config.ReadTimeout*time.Millisecond,
					config.WriteTimeout*time.Millisecond,
				)
				if err != nil {
					return nil, errors.Wrap(err, "redis dial error")
				}
				return conn, nil
			},
		}
	}
	return c
}

func (c Codis) GetConn() redis.Conn {
	cnt := len(c)
	if cnt == 0 {
		return nil
	}
	return c[rand.Intn(cnt)].Get()
}

// DoCmd 执行redis命令
func (c Codis) DoCmd(cmd string, args ...interface{}) (interface{}, error) {
	conn := c.GetConn()
	if conn == nil {
		return nil, errors.New("no redis connection")
	}
	defer conn.Close()

	return conn.Do(cmd, args...)
}

// DoScript 执行redis脚本
func (c Codis) DoScript(script *redis.Script, args ...interface{}) (interface{}, error) {
	conn := c.GetConn()
	if conn == nil {
		return nil, errors.New("no redis connection")
	}
	defer conn.Close()

	return script.Do(conn, args...)
}

// Close 关闭redis连接池
func (c Codis) Close() {
	for _, p := range c {
		p.Close()
	}
}

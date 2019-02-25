package codis

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

var config = Config{
	Addrs:          []string{"127.0.0.1:6379"},
	MaxIdle:        20,
	MaxActive:      50,
	ConnectTimeout: 3000,
	ReadTimeout:    3000,
	WriteTimeout:   3000,
	IdleTimeout:    30000,
	Wait:           true,
}

func TestCodis(t *testing.T) {
	co := New(config)
	defer co.Close()

	if _, err := co.DoCmd("SET", "KEY1", "20001"); err != nil {
		t.Fatal(err)
	}

	v, err := redis.Int(co.DoCmd("GET", "KEY1"))
	if err != nil {
		t.Fatal(err)
	}

	if v != 20001 {
		t.Fatal("Not Equal", v, 20001)
	}
}

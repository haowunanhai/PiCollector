package codis

import (
	"time"
)

// Config codis配置
type Config struct {
	Addrs          []string      `json:"addrs"`
	MaxIdle        int           `json:"max_idle"`
	MaxActive      int           `json:"max_active"`
	IdleTimeout    time.Duration `json:"idle_timeout"`
	ConnectTimeout time.Duration `json:"connect_timeout"`
	ReadTimeout    time.Duration `json:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"`
	Wait           bool          `json:"wait"`
	Db             int           `json:"db"`
}

package env

import (
	"flag"

	"common/log"
)

var (
	// 全局配置结构
	GlobalConfig *Config
	// 上报主机地址
	Host string
	// 服务端口号
	Port uint
	// Monitor端口号
	MonitorPort uint
)

func init() {
	// 设置命令行参数
	configFileName := flag.String("config", "conf/config.yaml", "config filename")
	//useIp := flag.Int("listen", 0, "listen host")
	//port := flag.Uint("port", 3608, "port")
	//monitorPort := flag.Uint("monitor", 4608, "monitor port")
	flag.Parse()

	//Port = *port
	//MonitorPort = *monitorPort

	var err error
	// 加载配置文件
	GlobalConfig, err = LoadConfig(*configFileName)
	if err != nil {
		log.Debug("load config file error %s %s\n", *configFileName, err)
		// 加载配置文件失败,生成空config结构
		GlobalConfig = &Config{}
	}

	// 设置日志
	if GlobalConfig.LogFile != "" {
		log.SetFile(GlobalConfig.LogFile)
	}
	// 设置日志级别
	if GlobalConfig.LogLevel != "" {
		log.SetLevel(GlobalConfig.LogLevel)
	}

}

package env

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config 的配置信息
type Config struct {
	LogFile  string                            `yaml:"logfile"`
	LogLevel string                            `yaml:"loglevel"`
	Data     map[string]map[string]interface{} `yaml:"data"`
}

// LoadConfig 装载配置文件内容
func LoadConfig(configFileName string) (*Config, error) {
	var config Config

	// 装载配置文件内容
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return &config, err
	}
	// 解析配置文件内容
	if err := yaml.Unmarshal(content, &config); err != nil {
		return &config, err
	}
	return &config, nil
}

func GetRedisAddrFromConfig(data map[string]interface{}, idc string) (string, int, string, error) {
	redisAddr, ok := data[idc].(string)
	if !ok {
		return redisAddr, 0, "", errors.New("redis addr not in conf")
	}

	redisDB, ok := data["db"].(int)
	if !ok {
		return redisAddr, redisDB, "", errors.New("redis db not in conf")
	}

	redisKey, _ := data["key"].(string)

	return redisAddr, redisDB, redisKey, nil
}

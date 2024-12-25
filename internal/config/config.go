package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

var (
	isDEV   bool
	conf    = Config{}
	name, _ = os.Hostname()
)

const (
	// SERVICE 服务名称
	SERVICE = "Rouguelike"
)

func init() {
	if strings.ToLower(GetEnv()) != "stable" && strings.ToLower(GetEnv()) != "prod" {
		isDEV = true
	}
}

// Config 配置
type Config struct {
	AppName  string   `toml:"appName"`
	Env      string   `toml:"env"`
	Timezone Timezone `toml:"timezone"`
	Width    int      `toml:"width"`
	Height   int      `toml:"height"`
	Tps      int      `toml:"tps"`
}

// GetLogEnvironment 格式化日志环境变量
func (c Config) GetLogEnvironment() string {
	switch c.Env {
	case "stable", "stable-internal":
		return "stable"
	}
	return strings.ReplaceAll(c.Env, "-", "_")
}

// LoadFile 从toml文件载入配置
func LoadFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithStack(err)
	}

	return Load(data)
}

// Load 从toml数据载入配置
func Load(b []byte) error {
	err := toml.Unmarshal(b, &conf)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// CheckConfig 检查配置有效性, 确保在程序启动阶段调用
func CheckConfig(conf *Config) error {
	if conf.Timezone == "" {
		conf.Timezone = "Asia/Shanghai"
	}
	return conf.Timezone.Check()
}

// GetConfig 获得配置对象
func GetConfig() Config {
	return conf
}

// GetEnv 获取环境名
// 环境名用来区分是否是生产环境, 另一方面用来记录日志,它会被写入到日志的e字段中
func GetEnv() string {
	return os.Getenv("ENV")
}

// IsDEV 是否开发环境
func IsDEV() bool {
	return isDEV
}

// IsAudit 是否是审核环境
func IsAudit() bool {
	return conf.Env == "audit"
}

// ServerName 获取服务器节点名称
func ServerName() string {
	return name
}

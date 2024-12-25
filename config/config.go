package config

import (
	_ "embed"
	"flag"
	"os"

	"github.com/pkg/errors"

	"github.com/tanqiangyes/rouguelike/internal/config"
	"github.com/tanqiangyes/rouguelike/pkg/logger"
)

//go:embed dev-config.toml
var devConf []byte

//go:embed i18n.toml
var langData []byte

// env 环境配置
var env string

// con 配置文件路径
var con string

// InitVar 设置启动参数
func InitVar() {
	flag.StringVar(&con, "config", "", "配置文件路径")
	flag.StringVar(&env, "env", "", "环境配置 stable = 正式环境")
	flag.Parse()

	if con != "" {
		if err := InitConfigFromPath(con); err != nil {
			panic(err)
		}
		return
	}

	if env != "" {
		if err := InitConfData(env); err != nil {
			panic(err)
		}
	}

}

// InitConfigFromPath 初始化配置文件
func InitConfigFromPath(conf string) error {
	if err := config.LoadFile(conf); err != nil {
		logger.NewMainEntry().WithError(err).Fatal("load config")
	}

	cfg := config.GetConfig()
	if err := config.CheckConfig(&cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// InitConfData 获取配置内容
func InitConfData(envs ...string) error {
	var env = os.Getenv("ENV")
	if len(envs) > 0 && envs[0] != "" {
		env = envs[0]
	}

	conf := devConf

	switch env {
	case "dev":
		conf = devConf
	default:
		panic("未定义的环境配置文件")
	}

	if err := config.Load(conf); err != nil {
		logger.NewEntry().WithError(err).Fatal("load config")
	}

	cfg := config.GetConfig()
	if err := config.CheckConfig(&cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GetLangData 获取多语言配置数据
func GetLangData() []byte {
	return langData
}

package config

import (
	"time"
)

// Timezone 时区对象
// TOML配置示例:
// timezone = "Asia/Shanghai"
type Timezone string

// Check 检查是否是有效时区
func (tz Timezone) Check() error {
	loc, err := time.LoadLocation(string(tz))
	if err == nil {
		time.Local = loc
	}
	return err
}

// Location 获得时区Location对象, 默认时区是合法时区,所以不会返回错误.
// 在程序启动阶段需要自行保证调用Check方法检查有效性
func (tz Timezone) Location() *time.Location {
	loc, _ := time.LoadLocation(string(tz))
	return loc
}

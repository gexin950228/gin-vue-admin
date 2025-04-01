package config

import (
	"gorm.io/gorm/logger"
	"strings"
)

type DsnProvider interface {
	Dsn() string
}

type GeneralDB struct {
	Prefix       string `yaml:"prefix" json:"prefix" mapstructure:"prefix"`                         // 数据库前缀
	Port         string `yaml:"port" json:"port" mapstructure:"port"`                               // 数据库端口
	Config       string `yaml:"config" json:"config" mapstructure:"config"`                         // 高级配置
	Dbname       string `yaml:"dbname" json:"dbname" mapstructure:"dbname"`                         // 数据库名
	Username     string `yaml:"username" json:"username" mapstructure:"username"`                   // 数据库账号
	Password     string `yaml:"password" json:"password" mapstructure:"password"`                   // 数据库密码
	Path         string `yaml:"path" json:"path" mapstructure:"path"`                               // 数据库地址
	Engine       string `yaml:"engine" json:"engine" mapstructure:"engine"`                         // 数据库引擎，默认InnoDB
	LogMode      string `yaml:"log_mode" json:"log_mode" mapstructure:"log_mode"`                   // 是否开启Gorm全局日志
	MaxIdleConns int    `yaml:"max-idle-conns" json:"max-idle-conns" mapstructure:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `yaml:"maxOpenConns" json:"max-open-conns" mapstructure:"max-open-conns"`   // 打开到数据库的最大连接数
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true表示开启
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (c GeneralDB) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

type SpecializedDB struct {
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	Disable   bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}

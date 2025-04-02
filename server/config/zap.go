package config

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Zap struct {
	Level         string `yaml:"level" json:"level" mapstructure:"level"`                            // 级别
	Prefix        string `yaml:"prefix" json:"prefix" mapstructure:"prefix"`                         // 日志前缀
	Format        string `yaml:"format" json:"format" mapstructure:"format"`                         // 输出
	Director      string `yaml:"director" json:"director" mapstructure:"director"`                   // 日志文件夹
	EncodeLevel   string `yaml:"encode_level" json:"encode_level" mapstructure:"encode_level"`       // 编码级
	StacktraceKey string `yaml:"stacktrace-key" json:"stacktrace-key" mapstructure:"stacktrace-key"` // 栈名
	ShowLine      bool   `yaml:"show_line" json:"show_line" mapstructure:"show_line"`                // 显示行
	LogInConsole  bool   `yaml:"log-in-console" json:"log-in-console" mapstructure:"log-in-console"` // 输出控制台
	RetentionDay  int    `yaml:"retention-day" json:"retention-day" mapstructure:"retention-day"`    // 日志保留天数
}

func (z *Zap) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(z.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

func Encoder(c *Zap) zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    c.LevelEncoder(),
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

// LevelEncoder 根据 EncodeLevel 返回 zapcore.LevelEncoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap) LevelEncoder() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器（默认）
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

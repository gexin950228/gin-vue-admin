package config

type Server struct {
	JWT       JWT     `json:"jwt"  json:"jwt" yaml:"jwt"`
	Zap       Zap     `json:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `json:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `json:"redis-list" json:"redis-list" yaml:"redis-list"`
	Mongo     Mongo   `json:"mongo" json:"mongo" yaml:"mongo"`
	Email     Email   `mapstructure:"email" json:"email" yaml:"email"`
	System    System  `json:"system" yaml:"system" mapstructure:"system"`
	Captcha   Captcha `json:"captcha" yaml:"captcha" mapstructure:"captcha"`
	Mysql     Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql     Mssql   `json:"mssql" yaml:"mssql" mapstructure:"mssql"`
}

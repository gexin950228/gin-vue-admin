package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`
	From     string `mapstructure:"from" json:"from" yaml:"from"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`
}

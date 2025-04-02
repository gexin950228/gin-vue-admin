package config

type System struct {
	DbType        string `json:"db-type" mapstructure:"db-type" yaml:"db-type"`
	OssType       string `json:"oss-type" mapstructure:"oss-type" yaml:"oss-type"`
	RouterPrefix  string `json:"router-prefix" mapstructure:"router-prefix" yaml:"router-prefix"`
	Addr          string `json:"addr" mapstructure:"addr" yaml:"addr"`
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time"  json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`                   // 使用mongo
	UseStrictAuth bool   `mapstructure:"use-strict-auth" json:"use-strict-auth" yaml:"use-strict-auth"` // 使用树形角色分配模式
}

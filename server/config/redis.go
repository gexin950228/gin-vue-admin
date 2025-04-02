package config

type Redis struct {
	Name         string   `mapstructrue:"name" json:"name" yaml:"name"`
	Addr         string   `mapstructrue:"addr" json:"addr" yaml:"addr"`
	Password     string   `mapstructrue:"password" json:"password" yaml:"password"`
	DB           int      `mapstructrue:"db" json:"db" yaml:"db"`
	UseCluster   bool     `mapstructrue:"useCluster" json:"useCluster" yaml:"useCluster"`
	ClusterAddrs []string `mapstructrue:"clusterAddrs" json:"clusterAddrs" yaml:"clusterAddrs"`
}

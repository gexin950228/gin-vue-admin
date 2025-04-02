package config

import (
	"fmt"
	"strings"
)

type Mongo struct {
	Coll             string       `json:"coll" mapstructure:"coll" yaml:"coll"`
	Options          string       `json:"options" mapstructure:"options"`
	Database         string       `json:"database" mapstructure:"database" yaml:"database"`
	Username         string       `json:"username" mapstructure:"username" yaml:"username"`
	Password         string       `json:"password" mapstructure:"password" yaml:"password"`
	AuthSource       string       `json:"auth-score" mapstructure:"auth-score" yaml:"auth-score"`
	MinPoolSize      string       `json:"min-pool-size" mapstructure:"min-pool-size" yaml:"min-pool-size"`
	SocketTimeoutMS  string       `json:"socket-timeout-ms" mapstructure:"socket-timeout-ms" yaml:"socket-timeout-ms"`
	ConnectTimeoutMS string       `json:"connect-timeout-ms" mapstructure:"connect-timeout-ms" yaml:"connect-timeout-ms"`
	IsZap            bool         `json:"is_zap" mapstructure:"is-zap" yaml:"is-zap"`
	Hosts            []*MongoHost `json:"hosts" mapstructure:"hosts" yaml:"hosts"`
}

type MongoHost struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port string `json:"port" mapstructure:"port" yaml:"port"`
}

func (x *Mongo) Uri() string {
	length := len(x.Hosts)
	hosts := make([]string, length)
	for i := 0; i < length; i++ {
		if x.Hosts[i].Host != "" && x.Hosts[i].Port != "" {
			hosts = append(hosts, x.Hosts[i].Host+":"+x.Hosts[i].Port)
		}
	}
	if x.Options != "" {
		return fmt.Sprintf("mongodb://%s/%s?%s", strings.Join(hosts, ";"), x.Database, x.Options)
	}
	return fmt.Sprintf("mongodb://%s/%s?%s", strings.Join(hosts, ";"), x.Database, x.Options)
}

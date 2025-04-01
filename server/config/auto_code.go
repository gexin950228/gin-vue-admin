package config

import (
	"path/filepath"
	"strings"
)

type AutoCode struct {
	Web     string `yaml:"web" json:"web" mapstructure:"web"`
	Root    string `yaml:"root" json:"root" mapstructure:"root"`
	Server  string `yaml:"server" json:"server" mapstructure:"server"`
	Module  string `yaml:"module" json:"module" mapstructure:"module"`
	ApiPath string `yaml:"apiPath" json:"apiPath" mapstructure:"apiPath"`
}

func (a AutoCode) WebPath() string {
	webs := strings.Split(a.Web, "/")
	if len(webs) == 0 {
		webs = strings.Split(a.Root, "\\")
	}
	return filepath.Join(webs...)
}

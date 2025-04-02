package config

type JWT struct {
	SigningKey  string `json:"signing-key" mapstructure:"signing-key" yaml:"signing-key"`
	ExpiresTime string `json:"expires-time" mapstructure:"expires-time" yaml:"expires-time"`
	BufferTime  string `json:"buffer-time" mapstructure:"buffer-time" yaml:"buffer-time"`
	Issuer      string `json:"issuer" mapstructure:"issuer" yaml:"issuer"`
}

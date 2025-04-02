package config

type Captcha struct {
	KeyLong     string `json:"keyLong" mapstructure:"keyLong" yaml:"keyLong"`    // 验证码长度
	ImgWidth    int    `json:"imgWidth" mapstructure:"imgWidth" yaml:"imgWidth"` // 验证码宽度
	ImgHeight   int    `json:"imgHeight" mapstructure:"imgHeight" yaml:"imgHeight"`
	OpenCaptcha int    `json:"openCaptcha" mapstructure:"openCaptcha" yaml:"openCaptcha"`
}

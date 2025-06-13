package config

type JWT struct {
	Secret      string `mapstructure:"secret" json:"secret" yaml:"secret"` // jwt签名
	ExpiresTime int    `mapstructure:"ttl" json:"ttl" yaml:"ttl"`          // 过期时间
	Iss         string `mapstructure:"iss" json:"iss" yaml:"iss"`          // Iss
	Sub         string `mapstructure:"sub" json:"sub" yaml:"sub"`          // Sub
}

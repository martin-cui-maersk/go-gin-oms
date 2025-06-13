package config

type Server struct {
	Version string `mapstructure:"version" json:"version" yaml:"version"` // version
	Port    string `mapstructure:"port" json:"port" yaml:"port"`          // gin port
	MySQL   MySQL  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT     JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}

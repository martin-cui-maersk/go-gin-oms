package config

type MySQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	User     string `mapstructure:"user" json:"user" yaml:"user"`             // 用户
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Database string `mapstructure:"database" json:"database" yaml:"database"` // 数据库
}

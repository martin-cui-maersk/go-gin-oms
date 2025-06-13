package config

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Database int    `mapstructure:"database" json:"database" yaml:"database"` // 单实例模式下redis的哪个数据库
}

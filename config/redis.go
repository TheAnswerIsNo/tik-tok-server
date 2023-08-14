package config

type Redis struct {
	Host     string `mapstructure:"172.16.32.69" json:"host" yaml:"host"`
	Port     int    `mapstructure:"50531" json:"port" yaml:"port"`
	DB       int    `mapstructure:"1" json:"db" yaml:"db"`
	Password string `mapstructure:"XFAWVFpF" json:"password" yaml:"password"`
}

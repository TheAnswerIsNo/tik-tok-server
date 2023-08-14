package config

type Configuration struct {
	App      App   `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log   `mapstructure:"log" json:"log" yaml:"log"`
	Database DB    `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt   `mapstructure:"jwt" json:"jwt"`
	Redis    Redis `mapstructure:"redis" json:"redis"`
}

type App struct {
	Env  string `mapstructure:"dev" json:"env" yaml:"env"`
	Port int    `mapstructure:"9000" json:"port" yaml:"port"`
	Name string `mapstructure:"douyin" json:"name" yaml:"name"`
	Url  string `mapstructure:"http://127.0.0.1" json:"url" yaml:"url"`
}

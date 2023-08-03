package config

type Configuration struct {
	App      App   `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log   `mapstructure:"log" json:"log" yaml:"log"`
	Database DB    `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt   `mapstructure:"jwt" json:"jwt"`
	Redis    Redis `mapstructure:"redis" json:"redis"`
}

type App struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Url  string `mapstructure:"url" json:"url" yaml:"url"`
}

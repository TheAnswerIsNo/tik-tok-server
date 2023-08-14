package config

type Log struct {
	Level      string `mapstructure:"info" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"./logs" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"app.log" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"json" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"true" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"3" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"500" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"28" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"true" json:"compress" yaml:"compress"`
}

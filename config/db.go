package config

type DB struct {
	Host                string `mapstructure:"172.16.32.69" json:"host" yaml:"host"`
	Port                int    `mapstructure:"52832" json:"port" yaml:"port"`
	Database            string `mapstructure:"tik_tok" json:"database" yaml:"database"`
	UserName            string `mapstructure:"root" json:"username" yaml:"username"`
	Password            string `mapstructure:"ddfpEHsK" json:"password" yaml:"password"`
	Charset             string `mapstructure:"utf8mb4" json:"charset" yaml:"charset"`
	ParseTime           bool   `mapstructure:"true" json:"parse_time" yaml:"parse_time"`
	Loc                 string `mapstructure:"Local" json:"loc" yaml:"loc"`
	MaxIdleConns        int    `mapstructure:"10" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"100" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"info" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"true" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"sql.log" json:"log_filename" yaml:"log_filename"`
}

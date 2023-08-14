package config

// Jwt token结构体
type Jwt struct {
	Secret                  string `mapstructure:"NoduBipzAWkVQzGGIlWGXf8p0QQfWZG1YG6sbF5aLg4z8N9BrF" json:"secret" yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"43200" json:"jwt_ttl" yaml:"jwt_ttl"`                                    // token 有效期（秒）
	JwtBlacklistGracePeriod int64  `mapstructure:"10" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"` // 黑名单宽限时间（秒）
	RefreshGracePeriod      int64  `mapstructure:"1800" json:"refresh_grace_period" yaml:"refresh_grace_period"`           // token 自动刷新宽限时间（秒）
}

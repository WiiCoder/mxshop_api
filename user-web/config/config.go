package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name" json:"name"`
	Host          string        `mapstructure:"host" json:"host"`
	Port          int           `mapstructure:"port" json:"port"`
	JWTInfo       JWTConfig     `mapstructure:"jwt" json:"jwt"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}

package config

type Configuration struct {
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL  MySQL  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PgSQL  PgSQL  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}

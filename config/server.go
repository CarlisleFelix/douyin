package config

type Server struct {
	Env        string `mapstructure:"env" json:"env" yaml:"env"`
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	ServerName string `mapstructure:"server_name" json:"server_name" yaml:"server_name"`
	ServerUrl  string `mapstructure:"server_url" json:"server_url" yaml:"server_url"`
	DbType     string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
}

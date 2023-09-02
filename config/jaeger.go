package config

type Jaeger struct {
	Endpoint    string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Environment string `mapstructure:"environment" json:"environment" yaml:"environment"`
	Servicename string `mapstructure:"servicename" json:"servicename" yaml:"servicename"`
	Id          int64  `mapstructure:"id" json:"id" yaml:"id"`
}

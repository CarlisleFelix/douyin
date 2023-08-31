package config

type RabbitMQ struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"` // rabbitmq用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // rabbitmq密码
}

func (m *RabbitMQ) Url() string {
	return "amqp://" + m.Username + ":" + m.Password + "@127.0.0.1:5672/"
}

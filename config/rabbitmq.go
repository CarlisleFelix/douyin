package config

// 消息队列
type RabbitMQ struct {
	Address     string // RabbitMQ地址
	Port        string // RabbitMQ端口
	UserName    string // RabbitMQ用户名
	Password    string // RabbitMQ密码
	VirtualHost string // RabbitMQ VirtualHost
}

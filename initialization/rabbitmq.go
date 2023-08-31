package initialization

import (
	"douyin/global"
	"douyin/middleware"
	"fmt"

	"github.com/streadway/amqp"
)

// li + chen
func InitTopic() {
	go middleware.FavoriteConsume()
	go middleware.VideoConsume()

	middleware.RmqFollowAdd = middleware.NewFollowRabbitMQ("follow_add")
	go middleware.RmqFollowAdd.Consumer()

	middleware.RmqFollowDel = middleware.NewFollowRabbitMQ("follow_del")
	go middleware.RmqFollowDel.Consumer()

	middleware.RmqCommentDel = middleware.NewCommentRabbitMQ("comment_del")
	go middleware.RmqCommentDel.Consumer()
}

func InitRabbitMQ() {
	url := global.SERVER_CONFIG.RabbitMQ.Url()
	dial, err := amqp.Dial(url)
	if err != nil {
		fmt.Println(url)
		global.SERVER_LOG.Fatal("rabbitmq connection err")
	}
	global.SERVER_RABBITMQ = dial
	InitTopic()
}

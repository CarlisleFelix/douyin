package middleware

import (
	"bytes"
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/service"
	"douyin/utils"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var RmqFollowAdd *FollowMQ
var RmqFollowDel *FollowMQ
var RmqCommentDel *CommentMQ

// chen
type FollowMQ struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

// chen
type CommentMQ struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

// chen
func NewFollowRabbitMQ(queueName string) *FollowMQ {
	followMQ := &FollowMQ{
		queueName: queueName,
	}

	cha, err := global.SERVER_RABBITMQ.Channel()
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq channel creation fail!")
	}
	followMQ.channel = cha
	return followMQ
}

// chen
func (f *FollowMQ) Consumer() {

	_, err := f.channel.QueueDeclare(f.queueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	//2、接收消息
	msgs, err := f.channel.Consume(
		f.queueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	switch f.queueName {
	case "follow_add":
		go f.consumerFollowAdd(msgs)
	case "follow_del":
		go f.consumerFollowDel(msgs)

	}

	<-forever
}

// chen
// 关系添加的消费方式。
func (f *FollowMQ) consumerFollowAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userId, _ := strconv.Atoi(params[0])
		targetId, _ := strconv.Atoi(params[1])
		// 日志记录。
		sql := fmt.Sprintf("CALL CreateFollowing(%v,%v)", targetId, userId)
		global.SERVER_LOG.Info("消费队列执行添加关系。SQL如下:", zap.String("sql", sql))
		// 执行SQL，注必须scan，该SQL才能被执行。
		if err := global.SERVER_DB.Raw(sql).Scan(nil).Error; nil != err {
			// 执行出错，打印日志。
			global.SERVER_LOG.Info(err.Error())
		}
	}
}

// chen
// 关系删除的消费方式。
func (f *FollowMQ) consumerFollowDel(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userId, _ := strconv.Atoi(params[0])
		targetId, _ := strconv.Atoi(params[1])
		// 日志记录。
		sql := fmt.Sprintf("CALL DeleteFollowing(%v,%v)", targetId, userId)
		//log.Printf("消费队列执行删除关系。SQL如下：%s", sql)
		// 执行SQL，注必须scan，该SQL才能被执行。
		if err := global.SERVER_DB.Raw(sql).Scan(nil).Error; nil != err {
			// 执行出错，打印日志。
			global.SERVER_LOG.Info(err.Error())
		}

	}
}

// NewCommentRabbitMQ 获取CommentMQ的对应队列。
func NewCommentRabbitMQ(queueName string) *CommentMQ {
	commentMQ := &CommentMQ{
		queueName: queueName,
	}

	cha, err := global.SERVER_RABBITMQ.Channel()
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq channel creation fail!")
	}
	commentMQ.channel = cha
	return commentMQ
}

// Consumer follow关系的消费逻辑。
func (c *CommentMQ) Consumer() {

	_, err := c.channel.QueueDeclare(c.queueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	//2、接收消息
	msg, err := c.channel.Consume(
		c.queueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	//只有删除逻辑
	forever := make(chan bool)
	go c.consumerCommentDel(msg)

	//log.Printf("[*] Waiting for messages,To exit press CTRL+C")

	<-forever
}

// 数据库中评论删除的消费方式。
func (c *CommentMQ) consumerCommentDel(msg <-chan amqp.Delivery) {
	for d := range msg {
		// 参数解析，只有一个评论id
		cId := fmt.Sprintf("%s", d.Body)
		commentId, _ := strconv.Atoi(cId)
		//log.Println("commentId:", commentId)
		//删除数据库中评论信息
		err := dao.DeleteCommentById(int64(commentId))
		if err != nil {
			log.Println(err)
		}
	}
}

// Publish Comment的发布配置。
func (c *CommentMQ) CommentPublish(message string) {

	_, err := c.channel.QueueDeclare(
		c.queueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		panic(err)
	}

	err1 := c.channel.Publish(
		c.exchange,
		c.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err1 != nil {
		panic(err)
	}
}

// Publish follow关系的发布配置。
func (f *FollowMQ) RelationPublish(message string) {

	_, err := f.channel.QueueDeclare(
		f.queueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		panic(err)
	}

	f.channel.Publish(
		f.exchange,
		f.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

// li
type FavoriteStruct struct {
	userId     int64
	videoId    int64
	actionType int64
}

// li
func FavoriteConsume() {
	//连接rabbitmq server

	//创建一个通道
	ch, err := global.SERVER_RABBITMQ.Channel()
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq channel creation fail!")
	}
	defer ch.Close()

	//声明队列
	q, err := ch.QueueDeclare(
		"favorite", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq queue creation fail!")
	}

	msgs, err := ch.Consume( // 注册一个消费者（接收消息）
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq consumer creation fail!")
	}
	forever := make(chan bool)
	favoriteStruct := FavoriteStruct{}
	// log.Println("favorit")
	go func() {
		for d := range msgs {
			binary.Read(bytes.NewReader(d.Body), binary.LittleEndian, &favoriteStruct)
			service.FavoriteAction(favoriteStruct.userId, strconv.FormatInt(favoriteStruct.videoId, 10), int32(favoriteStruct.actionType))
			// log.Println("消息处理成功")
		}
	}()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// li
func VideoConsume() {
	//连接rabbitmq server

	//创建一个通道
	ch, err := global.SERVER_RABBITMQ.Channel()
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq channel creation fail!")
	}
	defer ch.Close()

	//声明队列
	q, err := ch.QueueDeclare(
		"video", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq queue creation fail!")
	}

	msgs, err := ch.Consume( // 注册一个消费者（接收消息）
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		global.SERVER_LOG.Fatal("rabbitmq consumer creation fail!")
	}
	forever := make(chan bool)
	file := multipart.FileHeader{}
	var userId int64
	var title string
	go func() {
		for d := range msgs {
			binary.Read(bytes.NewReader(d.Body[0:8]), binary.LittleEndian, &userId)
			binary.Read(bytes.NewReader(d.Body[8:32]), binary.LittleEndian, &title)
			binary.Read(bytes.NewReader(d.Body[32:]), binary.LittleEndian, &file)
			//获得文件名并存储到本地
			fileName := fmt.Sprintf("%d_%s", userId, title) //标识名字
			//fmt.Println("filename:%s", fileName)
			fileExt := filepath.Ext(file.Filename)
			//fmt.Println("fileExt:%s", fileExt)
			finalFilename := fileName + fileExt
			//fmt.Println("finalFilename:%s", finalFilename)
			saveFilepath := filepath.Join("../tmp/", finalFilename) //路径+文件名
			//fmt.Println("saveFilepath:%s", saveFilepath)

			// c.SaveUploadedFile(data, saveFilepath)
			// =>
			src, err := file.Open()
			if err != nil {
				global.SERVER_LOG.Warn("file open fail!")
			}
			defer src.Close()

			if err = os.MkdirAll(filepath.Dir(saveFilepath), 0750); err != nil {
				global.SERVER_LOG.Warn("mkdir fail!")
			}

			out, err := os.Create(saveFilepath)
			if err != nil {
				global.SERVER_LOG.Warn("file open fail!")
			}
			defer out.Close()

			_, err = io.Copy(out, src)
			log.Println(err)
			//删除本地文件
			defer func() {
				err = os.Remove(saveFilepath)
				if err != nil {
					global.SERVER_LOG.Warn("File Deletion fail!")
				}
			}()

			//完成对象存储、以及数据库表活动
			curTime := utils.CurrentTimeInt()
			err = service.PublishService(userId, title, fileExt, curTime)
			// if err != nil {
			// 	c.JSON(http.StatusOK, response.Publish_Action_Response{
			// 		Response: response.Response{
			// 			StatusCode: 1,
			// 			StatusMsg:  err.Error(),
			// 		},
			// 	})
			// 	global.SERVER_LOG.Warn("Publish service fail!")
			// 	return
			// }
			global.SERVER_LOG.Info("video request Success!")
		}
	}()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// li
func FavoritePublish(user_id int64, video_id string, action_type int64, ctx context.Context) error {

	ctx, span := global.SERVER_VIDEO_TRACER.Start(ctx, "favorite mq service")
	defer span.End()

	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		global.SERVER_LOG.Fatal("fail to parse")
	}
	body := FavoriteStruct{
		userId:     user_id,
		videoId:    videoId,
		actionType: action_type,
	}
	//链接rabbitmq server
	//创建一个通道，用于完成任务
	ch, err := global.SERVER_RABBITMQ.Channel()
	if err != nil {
		global.SERVER_LOG.Fatal("Failed to open a channel")
	}
	//声明一个队列，用于消息发送
	q, err := ch.QueueDeclare(
		"favorite", // 队列名
		false,      // 持久化
		false,      // 是否自动删除
		false,      // 排他性
		false,      // no-wait
		nil,        // 附属参数
	)
	if err != nil {
		global.SERVER_LOG.Fatal("Failed to declare a queue")
	}
	buffer := make([]byte, 0, 24)
	binary.LittleEndian.PutUint64(buffer[0:8], uint64(body.userId))
	binary.LittleEndian.PutUint64(buffer[8:16], uint64(body.videoId))
	binary.LittleEndian.PutUint64(buffer[16:24], uint64(body.actionType))
	err = ch.Publish( // 发送消息（生产者）
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        buffer,
		})
	if err != nil {
		global.SERVER_LOG.Fatal("Failed to publish a message")
	}
	// log.Println("消息发送成功")
	return err
}

func VideoPublish(videoData []byte, ctx context.Context) error {

	ctx, span := global.SERVER_VIDEO_TRACER.Start(ctx, "video mq service")
	defer span.End()

	//链接rabbitmq server
	//创建一个通道，用于完成任务
	ch, err := global.SERVER_RABBITMQ.Channel()
	if err != nil {
		global.SERVER_LOG.Fatal("Failed to open a channel")
	}
	defer ch.Close()
	//声明一个队列，用于消息发送
	q, err := ch.QueueDeclare(
		"video", // 队列名
		false,   // 持久化
		false,   // 是否自动删除
		false,   // 排他性
		false,   // no-wait
		nil,     // 附属参数
	)
	if err != nil {
		global.SERVER_LOG.Fatal("Failed to declare a queue")
	}
	err = ch.Publish( // 发送消息（生产者）
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        videoData,
		})
	if err != nil {
		global.SERVER_LOG.Fatal("Failed to publish a message")
	}
	// log.Println("消息发送成功")
	return err
}

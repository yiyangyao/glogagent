package main

import (
	"fmt"
	"glogagent/config"
	"glogagent/kafka"
	"time"

	"github.com/Shopify/sarama"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// TODO: 1) 收集指定目录下的日志文件，发送到 kafka

func main() {
	var config0 = new(config.Config)
	// 0. 读取配置文件
	err := ini.MapTo(config0, "./config/config.ini")
	if err != nil {
		logrus.Errorf("load config err: %v", err)
		return
	}
	logrus.Infof("kafka addr: %+v", config0)
	// 1. 初始化
	kfkProducer, err := kafka.InitSyncProducer(config0.KafkaConfig)
	if err != nil {
		logrus.Errorf("kafka producer init err: %v", err)
		return
	}
	defer kfkProducer.Close()
	logrus.Info("init kafka producer init success")
	// 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "glogagent-demo"
	msg.Value = sarama.StringEncoder("kafka msg...")

	// 发送消息
	pid, offset, err := kfkProducer.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg error: %v", err)
		return
	}
	fmt.Printf("partionId: %v, offset: %v", pid, offset)
	time.Sleep(time.Second)
	// 2. 根据配置中的日志路径使用tail去收集日志
	// 3. 把日志通过sarama发送kafka
}

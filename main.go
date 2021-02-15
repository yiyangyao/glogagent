package main

import (
	"fmt"
	"glogagent/config"
	"glogagent/kafka"
	"glogagent/tailer"
	"time"

	"github.com/hpcloud/tail"

	"github.com/Shopify/sarama"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// TODO: 收集指定目录下的日志文件，发送到 kafka
func main() {
	var config0 = new(config.Config)
	// 0. 读取配置文件
	err := ini.MapTo(config0, "./config/config.ini")
	if err != nil {
		logrus.Errorf("load config err: %v", err)
		return
	}
	logrus.Infof("kafka addr: %+v", config0)
	// 1. 初始化kafka
	kfkProducer, err := kafka.InitSyncProducer(config0.KafkaConfig)
	if err != nil {
		logrus.Errorf("kafka producer init err: %v", err)
		return
	}
	defer kfkProducer.Close()
	logrus.Info("init kafka producer init success")

	// 2. 初始化 tailer
	tailer0, err := tailer.Init(config0.Collect)
	if err != nil {
		logrus.Errorf("tailer init err: %v", err)
		return
	}
	logrus.Info("tailer kafka producer init success")

	/// 3. 把日志通过sarama发送kafka
	var (
		msg *tail.Line
		ok  bool
	)

	for {
		msg, ok = <-tailer0.Lines
		if !ok {
			fmt.Printf("tailer file close reopen, filename: %s\n", tailer0.Filename)
			time.Sleep(time.Second)
			continue
		}
		pid, offset, err := kfkProducer.SendMessage(&sarama.ProducerMessage{
			Topic: "glogagent-demo",
			Value: sarama.StringEncoder(msg.Text),
		})
		if err != nil {
			fmt.Printf("send msg error: %v", err)
			return
		}
		fmt.Printf("partionId: %v, offset: %v", pid, offset)
	}
}

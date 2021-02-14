package kafka

import (
	"fmt"
	"glogagent/config"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func InitSyncProducer(c config.KafkaConfig) (sarama.SyncProducer, error) {
	// 1. 生产者配置
	kafConf := sarama.NewConfig()
	kafConf.Producer.RequiredAcks = sarama.WaitForAll
	kafConf.Producer.Partitioner = sarama.NewRandomPartitioner
	kafConf.Producer.Return.Successes = true

	// 2. 链接kafka
	addrs := strings.Split(c.Address, ",")
	client, err := sarama.NewSyncProducer(addrs, kafConf)
	if err != nil {
		logrus.Errorf("init kafka producer client err: %v", err)
		return nil, err
	}
	return client, nil
}

func main() {
	// 1. 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 2. 链接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("init producer client err: %v", err)
		return
	}
	defer client.Close()

	// 3. 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "glogagent-demo"
	msg.Value = sarama.StringEncoder("first msg...")

	// 4. 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg error: %v", err)
		return
	}
	fmt.Printf("partionId: %v, offset: %v", pid, offset)
}

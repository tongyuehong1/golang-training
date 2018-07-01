package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	w1, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w1.Ping()
	if err != nil {
		log.Fatal("should not be able to ping after Stop()")
		return
	}
	defer w1.Stop()
	topicName := "secondtest"
	msgCount := 2
	for i := 0; i < msgCount; i++ {
		err1 := w1.Publish(topicName, []byte("测试测试second"))
		if err1 != nil {
			log.Fatal("error")
		}
	}
}

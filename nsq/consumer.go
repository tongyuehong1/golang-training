package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
	"sync"
)

func main() {
	topicName := "secondtest"
	msgCount := 3
	for i := 0; i < msgCount; i++ {
		//time.Sleep(time.Millisecond * 20)
		go readMessage(topicName, i)
	}

	cleanup := make(chan os.Signal)
	signal.Notify(cleanup, os.Interrupt)
	fmt.Println("server is running....")

	quit := make(chan bool)
	go func() {
		select {
		case <- cleanup:
			fmt.Println("Received an interrupt , stoping service ...")
			for _, ele := range consumers {
				ele.StopChan <- 1
				ele.Stop()
			}
			quit <- true
		}
	}()
	<-quit
	fmt.Println("Shutdown server....")
}

type ConsumerHandle struct {
	q       *nsq.Consumer
	msgGood int
}

var consumers []*nsq.Consumer = make([]*nsq.Consumer, 0)
var mux *sync.Mutex = &sync.Mutex{}

func (h *ConsumerHandle) HandleMessage(message *nsq.Message) error {
	msg := string(message.Body) + "  " + strconv.Itoa(h.msgGood)
	fmt.Println(msg)

	return nil
}

func readMessage(topicName string, msgCount int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()

	config := nsq.NewConfig()
	config.MaxInFlight = 1000
	config.MaxBackoffDuration = 5 * time.Second

	//q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount), config)
	//q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount) + "#ephemeral", config)
	q, _ := nsq.NewConsumer(topicName, "ch"+strconv.Itoa(msgCount), config)

	h := &ConsumerHandle{q: q, msgGood: msgCount}
	q.AddHandler(h)

	//err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		fmt.Println("conect nsqd error")
		log.Println(err)
	}
	mux.Lock()
	consumers = append(consumers, q)
	mux.Unlock()
	<-q.StopChan
	fmt.Println("end....")
}

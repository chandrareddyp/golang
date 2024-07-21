package main

import (
	"fmt"
	"sync"
	"time"
)

const MaxMessageChannelSize = 2

/*
its basedon the article:
https://hackernoon.com/understanding-synccond-in-go-a-guide-for-beginners
*/
func main() {
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)

	var wg sync.WaitGroup

	messageChannel := NewMessageChannel(MaxMessageChannelSize)
	producer := NewProducer(messageChannel, cond)
	consumer := NewConsumer(messageChannel, cond)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=0;i<10;i++{
			producer.Produce(fmt.Sprintf("Message %d", i))
		}
	}()
	go func(){
		defer wg.Done()
		for i:=0;i<10;i++{
			consumer.Consume()
		}
	}()
	wg.Wait()
}

type MessageChannel struct {
	messages []string
	size     int
}

func NewMessageChannel(size int) *MessageChannel {
	return &MessageChannel{
		messages: make([]string, size),
		size:     size,
	}
}

func (mc *MessageChannel) AddMessage(msg string) {
	mc.messages = append(mc.messages, msg)
}
func (mc *MessageChannel) RemoveMessage() string {
	msg := mc.messages[0]
	mc.messages = mc.messages[1:]
	return msg
}
func (mc *MessageChannel) IsEmpty() bool {
	return len(mc.messages) == 0
}
func (mc *MessageChannel) IsFull() bool {
	return len(mc.messages) == mc.size
}

type Producer struct {
	messageChannel *MessageChannel
	cond           *sync.Cond
}

func NewProducer(mc *MessageChannel, cond *sync.Cond) *Producer{
	return &Producer{
		messageChannel : mc,
		cond: cond,
	}
}

func (p *Producer) Produce(msg string) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for p.messageChannel.IsFull() {
		fmt.Println("Producer waiting for consumer to consume")
		p.cond.Wait()
	}
	time.Sleep(1 * time.Second)
	p.messageChannel.AddMessage(msg)
	fmt.Println("Producer produced message:", msg)
	p.cond.Signal()
}

type Consumer struct {
	messageChannel *MessageChannel
	cond *sync.Cond
}

func NewConsumer(mc *MessageChannel, cond *sync.Cond) *Consumer {
	return &Consumer{
		messageChannel: mc,
		cond: cond,
	}
}

func (c *Consumer) Consume(){
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	for c.messageChannel.IsEmpty(){
		fmt.Println("Consumer waiting for producer to produce")
		c.cond.Wait()
	}
	msg:= c.messageChannel.RemoveMessage()
	fmt.Println("Consumer consumed message:",msg)
	c.cond.Signal()
}
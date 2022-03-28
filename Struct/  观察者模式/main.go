package main

import (
	"errors"
	"fmt"
	"sync"
)

type IClient interface {
	Name() string
	ConsumeMessage(string)
}

type BaseClient struct {
	name string
}

func NewBaseClient(name string) IClient {
	return &BaseClient{
		name: name,
	}
}
func (C *BaseClient) Name() string {
	return C.name
}

func (C *BaseClient) ConsumeMessage(msg string) {
	fmt.Printf("Client %s:%s\n", C.name, msg)
}

type Topic struct {
	Name    string
	clients map[string]IClient
	sync.RWMutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name:    name,
		clients: make(map[string]IClient),
		RWMutex: sync.RWMutex{},
	}
}
func (t *Topic) AddClient(c IClient) {
	if _, exist := t.clients[c.Name()]; !exist {
		t.Lock()
		t.clients[c.Name()] = c
		t.Unlock()
	}
}

func (t *Topic) Notify(msg string) {
	for _, client := range t.clients {
		client.ConsumeMessage(msg)
	}
}

type Server struct {
	Topics map[string]*Topic
	sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		Topics:  make(map[string]*Topic),
		RWMutex: sync.RWMutex{},
	}
}

func (S *Server) Subscribe(topicName string, c IClient) {
	S.Lock()
	topic, exist := S.Topics[topicName]
	S.Unlock()
	if exist {
		topic.AddClient(c)
		return
	}
	topic = NewTopic(topicName)
	topic.AddClient(c)
	S.Lock()
	S.Topics[topicName] = topic
	S.Unlock()
}

func (S *Server) Publish(topicName string, message string) error {
	S.Lock()
	topic, exist := S.Topics[topicName]
	S.Unlock()
	if exist {
		topic.Notify(message)
		return nil
	}
	return errors.New("当前频道无人定义")

}

func main() {
	server := NewServer()
	server.Subscribe("cctv1", NewBaseClient("河南"))

	server.Subscribe("cctv2", NewBaseClient("河南"))
	server.Subscribe("cctv2", NewBaseClient("河北"))

	fmt.Println("**********************************")
	server.Publish("cctv1", "天气预报")
	fmt.Println("**********************************")
	server.Publish("cctv2", "仙剑奇侠传3")
	fmt.Println("**********************************")
	server.Publish("cctv3", "喜羊羊")
	fmt.Println("**********************************")
}

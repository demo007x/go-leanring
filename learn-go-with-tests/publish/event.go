package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data interface{}
	Topic string
}
// DateChannel 是一个能接受DataEvent的channel
type DataChannel chan DataEvent
// DataChannelSlice 是一个能包含dataChannel的数据切片
type DataChannelSlice []DataChannel
// eventBus 存储有关订阅者感兴趣的特定主题信息
// EventBus 有subscribers 这是一个包含DataChannelSlice 的map。我们使用互斥锁来保护并发访问的读写
// 通过使用map和定义topics， 他应许我们轻松的组织事件，主题被视为map的键，当有人发布时我们可以通过键轻松的找到了主题，然后将事件传播到channel中已进行进一步的处理。
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm sync.RWMutex
}

//订阅主题，使用channel， 他就行传统方法中的回调一样，当发布者向主题发布数据时，channel将接受数据
func (eb *EventBus) Subscribe(topic string, ch DataChannel)  {
	eb.rm.Lock()

	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}

	eb.rm.Unlock()
}
// 简单的说，我们将订阅这添加到channel中， 切片人后将给结构枷锁，最后在操作完成后将其解锁
// 发布主题
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.Lock()
	if chans, found := eb.subscribers[topic]; found {
		// 这样做的原因是因为切片引用相同的数组，即使他们是安置传递的
		// 因此我们正在使用我们的元素创建一个新的切片
		channels := append(DataChannelSlice{}, chans...)
		 go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			 for _, ch := range dataChannelSlices {
				 ch <- data
			 }
		 }(DataEvent{Data:data, Topic:topic}, channels)
	}

	eb.rm.Unlock()
}

// 创建一个时间总线的实例
var eb  = &EventBus{
	subscribers: map[string]DataChannelSlice{},
	rm:          sync.RWMutex{},
}
//为了测试新创建的事件总线，我们将创建一个以随机间隔时间发布到指定主题的方法。

func publisTo(topic string, data string)  {
	// 死循环， 监听订阅者的发布
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	}
}
//接下来，我们需要一个可以收听主题的 main 函数。它使用辅助方法打印出事件的数据。
func printDataEvent(ch string, data DataEvent)  {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic3", ch3)

	go publisTo("topic1", "Hi topic 1")
	go publisTo("topic2", "Welcome to topic 2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)

		}
	}
}



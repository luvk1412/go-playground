package main

import (
	"event_orchestration/event"
	"fmt"
	"math/rand"
)

var channels = make(map[string]chan event.Event)

func send(e event.Event) {
	consumerChannel, exists := channels[e.EType]
	if exists {
		consumerChannel <- e
	}
}

func registerConsumer(consumer *event.Consumer) {
	ch := make(chan event.Event)
	channels[consumer.EType] = ch
	consumer.Start(ch)
}

func main() {
	eTypes := []string{"a", "b", "c"}
	for _, eType := range eTypes {
		registerConsumer(&event.Consumer{EType: eType})
	}

	eventCt := 10
	producer := event.Producer{}

	for i := 0; i < eventCt; i++ {
		event := producer.Produce(eTypes[rand.Intn(len(eTypes))])
		send(event)
		fmt.Println("Sent", event)
	}
}

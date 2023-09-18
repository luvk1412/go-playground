package event

import (
	"math/rand"
	"time"
)

type Producer struct {
}

func (producer *Producer) Produce(eType string) Event {
	return Event{rand.Intn(100), eType, time.Now().UnixMilli()}
}

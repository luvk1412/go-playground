package event

import (
	"fmt"
	"time"
)

type Consumer struct {
	EType string
}

func (c *Consumer) consume(e Event) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Consumed", e, "at", time.Now().UnixMilli())
}

func (c *Consumer) Start(ch chan Event){
	go func(){
		for {
			c.consume(<-ch)
		}
	}()
}

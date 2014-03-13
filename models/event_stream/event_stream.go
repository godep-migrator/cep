package event_stream

import (
	"github.com/thresholderio/go-processing/models/context"
	"log"
	"sync"
)

type EventStream struct {
	Queue   [][]string
	Stop    chan bool
	Exit    chan bool
	Running bool
	*sync.WaitGroup
}

var EventStreams = make(map[string]EventStream)

func StartStreams() error {
	for _, e := range EventStreams {
		e.Add(1)
		go e.Watch()
	}

	return nil
}

func (self EventStream) Watch() {
	self.Running = true

	for {
		select {
		case <-self.Stop:
			self.Running = false
			self.Exit <- true
			break
		default:
			// Dequeue an event.
			if len(self.Queue) > 0 {
				tuple := self.Queue[0]
				log.Printf("received event: %+v\n", tuple)
				copy(self.Queue[0:], self.Queue[1:])
				self.Queue[len(self.Queue)-1] = nil
				self.Queue = self.Queue[:len(self.Queue)-1]

				context := &context.Context{}
				context.BuildUserContext(tuple[0], tuple[1])
				log.Println("UserContext: %+v\n", context.UserContext)
			}
		}
	}
}

func TeardownStreams() {
	log.Println("Cleaning up event streams...")
	for _, e := range EventStreams {
		e.Quit()
	}
}

func (self EventStream) Quit() {
	self.Stop <- true
	<-self.Exit
	self.Done()
}

package initializer

import (
	es "github.com/jeffchao/cep/models/event_stream"
	"sync"
)

func Initialize(eventStreams map[string]es.EventStream) {
	eventStream := es.EventStream{
		[][]string{[]string{"vx-1", "delay"}, []string{"vx-2", "delay"}},
		make(chan bool),
		make(chan bool),
		false,
		&sync.WaitGroup{},
	}
	eventStreams["flight-tracker"] = eventStream
}

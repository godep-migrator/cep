package main

import (
	"github.com/jeffchao/cep/config/cassandra"
	es "github.com/jeffchao/cep/models/event_stream"
	"github.com/jeffchao/cep/support/initializer"
	"github.com/jeffchao/cep/support/seeds"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	go HandleSignals()
	work()
}

func work() {
	es.StartStreams()
	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

// Catch signals that might terminate the process on behalf all goroutines.
func HandleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	for s := range signals {
		switch s {
		case syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			Quit()
		}
	}
}

// Ensure goroutines are cleaned up gracefully before exiting.
func Quit() {
	log.Println("Waiting for cleanup...")
	es.TeardownStreams()
	log.Println("Exiting")
	os.Exit(1)
}

func main() {
	log.Println("Starting processor")

	cassandra.CQL()
	defer cassandra.Session.Close()

	seeds.SeedUsers()
	seeds.SeedUserStates()
	seeds.SeedFlights()
	seeds.SeedUsersByFlight()
	initializer.Initialize(es.EventStreams)

	Run()
}

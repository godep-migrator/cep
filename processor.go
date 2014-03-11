package main

import (
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
	for {
		time.Sleep(1000 * time.Millisecond)
		log.Println("working...")
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
	log.Println("Exiting")
	os.Exit(1)
}

func main() {
	log.Println("Starting processor")
	Run()
}

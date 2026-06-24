package main

import (
	"log"
	"sync"

	"github.com/SiddJeph/email-dispatcher/internal/config"
	"github.com/SiddJeph/email-dispatcher/internal/email"
	"github.com/SiddJeph/email-dispatcher/internal/producer"
)

func main() {
	ch := make(chan email.Recipient)

	go func() {
		if err := producer.LoadRecipients(config.CSVPath, ch); err != nil {
			log.Fatalf("Failed to load recipients: %v", err)
		}
	}()

	var wg sync.WaitGroup
	for i := 1; i <= config.WorkerCount; i++ {
		wg.Add(1)
		go email.Worker(i, ch, &wg)
	}
	wg.Wait()
}

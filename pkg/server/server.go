package server

import (
	"sync"

	"github.com/linuxuser586/pki/pkg/server/rpc"
)

// Start the server
func Start() {
	log.Info("starting PKI servers...")
	messages := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		messages <- rpc.Start()
		messages <- ""
	}()
	go func() {
		for msg := range messages {
			log.Info(msg)
		}
	}()
	wg.Wait()
	log.Info("stopped PKI servers")
}

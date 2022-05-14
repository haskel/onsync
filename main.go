package main

import (
	"fmt"
	"log"
	"sync"
	config "ysync/configuration"
	s "ysync/syncer"
)

func main() {
	configuration, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Parse config error: %v", err)
	}

	wg := sync.WaitGroup{}

	for name, syncConfig := range configuration.Syncs {
		fmt.Printf("Start syncing %s\n", name)
		source := configuration.Targets[syncConfig.Source]
		target := configuration.Targets[syncConfig.Target]

		wg.Add(1)

		syncer := s.Syncer{}
		go syncer.Start(name, source, target, syncConfig, &wg)
	}

	wg.Wait()
}

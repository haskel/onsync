package syncer

import (
	"fmt"
	"github.com/briandowns/spinner"
	"path/filepath"
	"strings"
	"sync"
	"time"
	config "ysync/configuration"
	"ysync/transfer"
	"ysync/watcher"
)

type Syncer struct {
}

func (*Syncer) Start(
	name string,
	source config.Target,
	target config.Target,
	syncConfig config.Sync,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	if source.Host != config.Localhost {
		fmt.Printf("Only local sources supported %s\n", name)
		return
	}

	pathsChan := make(watcher.PathsChan, 1000)
	go watcher.Watch(pathsChan, source)

	go func(pathsChan chan string, source config.Target, target config.Target) {
		s := spinner.New(spinner.CharSets[13], 100*time.Millisecond)
		err := s.Color("yellow", "bold")
		if err != nil {
			return
		}
		s.Start()

		scp := transfer.Scp{}
		rsync := transfer.Rsync{}

		tick := time.Tick(3 * time.Second)
		syncJob := SyncJob{}

		go rsync.Sync([]string{}, source, target, syncConfig)

		for {
			select {
			case path := <-pathsChan:
				if isExcluded(path, syncConfig) {
					break
				}

				fmt.Printf("\rAdded path %s\n", path)
				syncJob.AddPath(path)

			case <-tick:
				//fmt.Println(paths)
				//sync := strategy.GetSyncer(&syncJob)
				switch true {
				case syncJob.Length() == 1:
					var path, err = syncJob.GetFirst()
					if err != nil {
						break
					}
					go scp.Sync(path, source, target)
					break

				case syncJob.Length() > 1:
					go rsync.Sync(syncJob.GetPaths(), source, target, syncConfig)
					break
				}

				syncJob.Clear()
			}
		}

	}(pathsChan, source, target)

	select {}

}

func isExcluded(path string, syncConfig config.Sync) bool {
	for _, pattern := range syncConfig.Directories.Excluded {
		if strings.HasSuffix(path, "~") { // todo: move to rules
			return true
		}

		if strings.HasPrefix(path, pattern) {
			return true
		}

		match, err := filepath.Match(pattern, path)
		if err != nil || match {
			return true
		}
	}

	return false
}

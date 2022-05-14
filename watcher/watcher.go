package watcher

import (
	"fmt"
	"github.com/fsnotify/fsevents"
	"log"
	"strings"
	"time"
	config "ysync/configuration"
)

var noteDescription = map[fsevents.EventFlags]string{
	fsevents.MustScanSubDirs: "MustScanSubdirs",
	fsevents.UserDropped:     "UserDropped",
	fsevents.KernelDropped:   "KernelDropped",
	fsevents.EventIDsWrapped: "EventIDsWrapped",
	fsevents.HistoryDone:     "HistoryDone",
	fsevents.RootChanged:     "RootChanged",
	fsevents.Mount:           "Mount",
	fsevents.Unmount:         "Unmount",

	fsevents.ItemCreated:       "Created",
	fsevents.ItemRemoved:       "Removed",
	fsevents.ItemInodeMetaMod:  "InodeMetaMod",
	fsevents.ItemRenamed:       "Renamed",
	fsevents.ItemModified:      "Modified",
	fsevents.ItemFinderInfoMod: "FinderInfoMod",
	fsevents.ItemChangeOwner:   "ChangeOwner",
	fsevents.ItemXattrMod:      "XAttrMod",
	fsevents.ItemIsFile:        "IsFile",
	fsevents.ItemIsDir:         "IsDir",
	fsevents.ItemIsSymlink:     "IsSymLink",
}

var WatchedEvents = fsevents.ItemCreated | fsevents.ItemRemoved | fsevents.ItemRenamed | fsevents.ItemModified

type Watcher struct {
}

func printEventNames() {
	for bit, description := range noteDescription {
		if bit&WatchedEvents == bit {
			fmt.Println(description)
		}
	}
}

type PathsChan chan string

func Watch(pathsChan PathsChan, source config.Target) {
	eventChan := initEventChan(source.Path)

	go func() {
		for msg := range eventChan {
			for _, event := range msg {
				if event.Flags&WatchedEvents <= 0 {
					continue
				}

				relPath := strings.TrimPrefix("/"+event.Path, source.Path+"/")
				//relPath, _ := filepath.Rel(source.Path, event.Path)
				//fmt.Printf("relPath: %s, source.Path: %s, event.Path: %s, event: %d\n", relPath, source.Path, event.Path, event.Flags)

				pathsChan <- relPath
			}
		}
	}()
}

func initEventChan(path string) chan []fsevents.Event {
	device, err := fsevents.DeviceForPath(path)
	if err != nil {
		log.Fatalf("Failed to retrieve device for path: %v", err)
	}

	//log.Print(device)
	//log.Println(fsevents.EventIDForDeviceBeforeTime(device, time.Now()))

	log.Printf("%s\n", path)

	es := &fsevents.EventStream{
		Paths:   []string{path},
		Latency: 500 * time.Millisecond,
		Device:  device,
		Flags:   fsevents.FileEvents | fsevents.WatchRoot,
	}
	es.Start()

	//log.Println("Device UUID", fsevents.GetDeviceUUID(device))

	return es.Events
}

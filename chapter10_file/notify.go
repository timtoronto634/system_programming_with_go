package chapter10file

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func Notify() {
	counter := 0
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					log.Println("created file:", event.Name)
					counter++
				} else if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					counter++
				} else if event.Has(fsnotify.Remove) {
					log.Println("removed file:", event.Name)
					counter++
				} else if event.Has(fsnotify.Rename) {
					log.Println("renamed file:", event.Name)
					counter++
				} else if event.Has(fsnotify.Chmod) {
					log.Println("changed permission file:", event.Name)
					counter++
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
			if counter > 3 {
				done <- true
			}
		}
	}()

	// Add a path.
	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

/*
@Time : 2019-07-25 14:52
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"log"
	"path/filepath"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()
	fmt.Println(filepath.Abs("./"))
	err = watcher.Watch("/Users/lx/Desktop/release/testm/")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()

}

/*
@Time : 2019-07-25 16:45
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"encoding/json"
	"github.com/fvbock/endless"
	"github.com/hashicorp/consul/logger"
	"github.com/howeyc/fsnotify"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func main() {
	app := iris.Default()
	app.Get("/", func(ctx iris.Context) {
		//log.Println("hot update ", ctx.Params().Get("id"))
		ctx.JSON(iris.Map{"code": 0, "msg": "testmain " + time.Now().String()})
	})
	app.Build()
	err := endless.ListenAndServe("localhost:8008", app.Router)
	//err := app.Run(iris.Addr("8008"))
	if err != nil {
		//panic(err.Error())
	}

}

var (
	restartChan = make(chan bool)
	started     bool
)

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				log.Println("event:", event)
				if event.IsModify() {
					log.Println("notify runner to do the ln -s and restart server.")
					restartChan <- true
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/path/to/current.conf")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan bool)
}

func run() {
	for {

		<-restartChan

		c, err := ioutil.ReadFile("/path/to/current.conf")
		if err != nil {
			log.Println("current.conf read error:", err)
			return
		}

		var j interface{}
		err = json.Unmarshal(c, &j)
		if err != nil {
			log.Println("current.conf parse error:", err)
			return
		}

		parsed, ok := j.(map[string]interface{})
		if !ok {
			log.Println("current.conf parse error: mapping errors")
			return
		}

		exec.Command("rm", "app.bin").Run()
		exec.Command("ln", "-s", parsed["bin.file"].(string), "app.bin").Run()

		exec.Command("rm", "app.conf").Run()
		exec.Command("ln", "-s", parsed["cfg.file"].(string), "app.cfg").Run()

		if !started {
			cmd := exec.Command("./app.bin", "-c", "app.cfg")
			started = true
		} else {
			processes, _ := ps.Processes()
			for _, v := range processes {
				if strings.Contains(v.Executable(), parsed["bin.file"]) {
					process, _ := os.FindProcess(v.Pid())
					process.Signal(syscall.SIGHUP)
				}
			}
		}
	}
}

func forkProcess() {

}

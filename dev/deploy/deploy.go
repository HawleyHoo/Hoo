/*
@Time : 2019-06-20 12:25
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func relaunching() {
	cmd := exec.Command("sh", "./deploy.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func restart(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>deploy server: restarting webserver...</h1>")
	relaunching()
	io.WriteString(w, "<h1>deploy server: webserver restarted!</h1>")
}

func main() {
	//relaunching()
	//return
	//fmt.Println("begin exec deploy.sh...")
	cmd := exec.Command("sh", "./deploy.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("err:", err)
	}
	log.Println(string(out))

	//err = cmd.Start()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = cmd.Wait()
	//fmt.Println("complete!")

	//http.HandleFunc("/", restart)
	//http.ListenAndServe(":5000", nil)
}

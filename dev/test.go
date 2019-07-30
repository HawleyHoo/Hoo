/*
@Time : 2019-06-26 10:59
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"github.com/gogf/gf/g/os/gproc"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"syscall"
	"time"
)

/*

stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(os.Stderr, "error=>", err.Error())
		return
	}

reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	go func() {
		for {
			//line, err2 := reader.ReadString('\n')
			line, isprefix, err2 := reader.ReadLine()
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println("read: ", line, isprefix)
			time.Sleep(time.Second * 2)
		}
	}()
*/
func testgf() {
	fmt.Println("main", os.Getpid(), os.Getppid())
	//go func() {
	//err := gproc.ShellRun("nohup ./testmain & echo $! > pidf.txt")
	out, err := gproc.ShellExec("/usr/bin/nohup ./testmain &")
	//out, err := gproc.ShellExec("ls -l")
	fmt.Println("err:", err, out)
	//}()

	select {}

}

func testls() {
	cmd := exec.Command("sh", "-c", "nohup ./testmain >> nohup.out &")
	cmd.Dir = "/Users/lx/Desktop/release/testm"
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	//cmd := exec.CommandContext(ctx,"nohup", "./main",  "&")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Start err:", err)
	}

	time.Sleep(time.Second * 2)
	os.Stdin.WriteString("\n")
	err = cmd.Wait()
	if err != nil {
		fmt.Println("wait err:", err)
	}
	fmt.Printf("cmd:%+v\n", cmd)
	fmt.Println(cmd.ProcessState.String(), "pid:", cmd.Process.Pid)
}

func main() {

	fmt.Println("main", os.Getpid(), os.Getppid())

	testls()

	return

	//c := "nohup ./testmain & echo $! > pidf.txt"
	//c := "nohup /root/go/trade/trade > myout.log 2>&1 &"

	c := "nohup ./testmain > myout.log 2>&1 &"
	//c := "setsid ./testmain &"
	cmd := exec.Command("sh", "-c", c)
	//cmd := exec.CommandContext(ctx,"nohup", "./main",  "&")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Start err:", err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("wait err:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("main server.")
		io.WriteString(w, "<h1>This is A main Index Page!  </h1>")
	})
	port := "8081"
	log.Println("listen server: localhost:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicln(err.Error())
	}

}

func test02() {
	u, _ := user.Current()
	//fmt.Println(os.Environ())
	fmt.Println(os.Getpid(), os.Getppid())
	fmt.Printf("user:%+v \n", u)

	pro, err := os.FindProcess(4054)
	if err != nil {
		fmt.Println("FindProcess", err.Error())
		return
	}
	fmt.Println("process info:", pro.Release())
	err = pro.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Println("Signal: ", err.Error())
		//err = startProcess("./main", nil)
		//go gproc.ShellRun(`nohup ./main & echo $! > pidf.txt`)

		cmd := exec.Command("sh", "./start.sh")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("sh:", string(out))

	}
	time.Sleep(time.Second * 10)
	fmt.Println("Signal SIGTERM:", pro.Pid)
}

//启动进程
func startProcess(exePath string, args []string) error {
	attr := &os.ProcAttr{
		//files指定新进程继承的活动文件对象
		//前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		//新进程的环境变量
		Env: os.Environ(),
	}

	p, err := os.StartProcess(exePath, args, attr)
	if err != nil {
		return err
	}
	fmt.Println(exePath, "进程启动")
	pstate, err := p.Wait()
	if err != nil {
		return err
	}
	fmt.Println("pstate:", pstate)
	return nil
}

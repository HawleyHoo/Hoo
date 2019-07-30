package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	//"code.google.com/p/go.net/websocket"
)

const tpl = `<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:1234";
        window.onload = function() {
            console.log("onload");
            sock = new WebSocket(wsuri);
            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }
            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }
            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
            }
        };
        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
        };
    </script>
    <h1>WebSocket Echo Test</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="Hello, world!">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>`

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "hehe :" + reply + strconv.Itoa(rand.Int())
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("tpl")
	content, _ := t.Parse(tpl)
	content.Execute(w, nil)
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.Handle("/send", http.HandlerFunc(myHandler))

	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8080/echo"

func maintest() {
	/*
		客户端使用websocket.Dial(url, “”, origin) 进行websocket连接，但是origin参数并没有实际调用。
		使用websocket进行数据的发送和接受。非常有意思的事情是，如果客户端和服务端都是用go写，用的都是websocket这个对象。
		函数调用都是一样的，只不过一个写一个读数据而已。
	*/
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("hello, world!你好")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])

	ws.Close() //关闭连接
}

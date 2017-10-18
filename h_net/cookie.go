package main

import (
	"net/http"
	"io"
)



func main() {
	http.HandleFunc("/", Cookie2)

	http.ListenAndServe(":9090", nil)

}

func Cookie(w http.ResponseWriter, r *http.Request)  {
	ck := http.Cookie{
		Name: "myCookie",
		Value: "hello",
		Path: "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	http.SetCookie(w, ck)
	ck2 , err := r.Cookie("myCookie")
	if er {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}


func Cookie2(w http.ResponseWriter, r *http.Request)  {
	ck := http.Cookie{
		Name: "myCookie",
		Value: "hello",
		Path: "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	w.Header().Set("Set-Cookie", ck.String())

	ck2 , err := r.Cookie("myCookie")
	if er {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}



/*type myHandler struct {

}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main()  {
	server := http.Server{
		Addr: ":9000",
		Handler: &myHandler{},
		ReadHeaderTimeout: 5 * time.Second,
	}




	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

	mux["/hello"] = sayHello
	mux["/yes"] = sayYes


	err := server.ListenAndServe()
	if err == nil {
		log.Fatal(err)
	}
}



func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 1")
}
func sayYes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Yes , this is version 1")
}

func (haha *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if h, ok := mux[r.URL.String()];ok {
		h(w, r)
		return
	}
	io.WriteString(w,"URL:" + r.URL.String())
}*/


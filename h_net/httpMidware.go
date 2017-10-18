package main

import "net/http"

type SingleHost struct {
	handler    http.Handler
	alloweHost string
}

func (this *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == this.alloweHost {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func main() {
	handler := SingleHost{
		handler:    http.HandlerFunc(myHandler),
		alloweHost: "example.com",
	}

	http.ListenAndServe(":9092", handler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

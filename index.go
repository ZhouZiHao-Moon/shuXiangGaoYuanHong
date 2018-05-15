package main

import (
	"io/ioutil"
	"net/http"
)

var html []byte

func file(w http.ResponseWriter, req *http.Request) {
	buff, err := ioutil.ReadFile(req.URL.Path[1:])
	if err != nil {
		buff = []byte("open file fail!")
	}
	w.Write(buff)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		buff, err := ioutil.ReadFile("index.html")
		if err != nil {
			buff = []byte("open file fail!")
		}
		w.Write(buff)
	}
}

func books(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("100"))
}

func donate(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("[[1,10,11,100],[2,20,22,200]]"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/static/", file)
	http.HandleFunc("/books/", books)
	http.HandleFunc("/donate/", donate)
	http.ListenAndServe(":14250", nil)
}

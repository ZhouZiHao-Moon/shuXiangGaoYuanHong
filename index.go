package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
)

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
	xlsx, err := excelize.OpenFile("books.xlsx")
	if err != nil {
		w.Write([]byte("open file fail!"))
		return
	}
	w.Write([]byte(xlsx.GetCellValue("Sheet1", "A1")))
}

func donate(w http.ResponseWriter, req *http.Request) {
	xlsx, err := excelize.OpenFile("books.xlsx")
	if err != nil {
		w.Write([]byte("open file fail!"))
		return
	}
	rows := xlsx.GetRows("Sheet1")
	w.Write([]byte("["))
	for key, row := range rows[2:] {
		if key != 0 {
			w.Write([]byte(","))
		}
		w.Write([]byte("["))
		for key2, colCell := range row {
			if key2 != 0 {
				w.Write([]byte(","))
			}
			w.Write([]byte(colCell))
		}
		w.Write([]byte("]"))
	}
	w.Write([]byte("]"))
}

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/", index)
	http.HandleFunc("/static/", file)
	http.HandleFunc("/books/", books)
	http.HandleFunc("/donate/", donate)
	http.ListenAndServe(":14250", nil)
}

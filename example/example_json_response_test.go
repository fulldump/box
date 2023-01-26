package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/fulldump/box"
)

type MyResponse struct {
	Name string
	Age  int
}

func Example_Json_responses() {

	b := box.NewBox()
	b.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) MyResponse {
		return MyResponse{
			Name: "Fulanez",
			Age:  33,
		}
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/hello")
	io.Copy(os.Stdout, resp.Body)

	// Output:
	// {"Name":"Fulanez","Age":33}
}

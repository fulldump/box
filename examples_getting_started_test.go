package box_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/fulldump/box"
)

func ExampleB_ServeHTTP_gettingStarted() {

	b := box.NewBox()
	b.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World!"))
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/hello")
	io.Copy(os.Stdout, resp.Body)

	// Output:
	// World!
}

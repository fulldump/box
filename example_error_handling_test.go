package box_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/fulldump/box"
)

func Example_Error_handling() {

	b := box.NewBox()

	b.Use(box.PrettyError)

	b.Handle("GET", "/articles", func() (*Article, error) {
		return nil, errors.New("could not connect to the database")
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/articles")
	io.Copy(os.Stdout, resp.Body)

	// Output:
	// could not connect to the database
}

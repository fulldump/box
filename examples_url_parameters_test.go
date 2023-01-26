package box_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/fulldump/box"
)

func Example_Url_parameters() {

	b := box.NewBox()
	b.Handle("GET", "/articles/{article-id}", func(w http.ResponseWriter, r *http.Request) string {
		articleID := box.Param(r, "article-id")
		return "ArticleID is " + articleID
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/articles/123")
	io.Copy(os.Stdout, resp.Body)

	// Output:
	// "ArticleID is 123"
}

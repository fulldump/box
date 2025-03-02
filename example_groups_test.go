package box_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/fulldump/box"
)

func ExampleR_Group() {

	b := box.NewBox()

	v0 := b.Group("/v0")
	v0.Use(box.SetResponseHeader("Content-Type", "application/json"))

	v0.Handle("GET", "/articles", ListArticles)
	v0.Handle("POST", "/articles", CreateArticle)

	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/v0/articles")
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	// Output:
	// Content-Type: application/json
}

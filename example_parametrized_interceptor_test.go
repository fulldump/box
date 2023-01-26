package box_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/fulldump/box"
)

func Example_ParametrizedInterceptor() {

	b := box.NewBox()

	b.Use(
		box.SetResponseHeader("Server", "My server name"),
		box.SetResponseHeader("Version", "v3.2.1"),
	)

	b.Handle("GET", "/articles/{article-id}", func(r *http.Request) *Article {
		return &Article{
			Id:    box.Param(r, "article-id"),
			Title: "Example",
		}
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/something")
	for _, key := range []string{"Server", "Version", "Content-Length"} {
		fmt.Println(key+":", resp.Header.Get(key))
	}
	// Output:
	// Server: My server name
	// Version: v3.2.1
	// Content-Length: 0
}

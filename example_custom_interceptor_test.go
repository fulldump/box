package box_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/fulldump/box"
)

func MyCustomInterceptor(next box.H) box.H {
	return func(ctx context.Context) {
		w := box.GetResponse(ctx)
		w.Header().Set("Server", "MyServer")
		next(ctx) // continue the flow
	}
}

func Example_CustomInterceptor() {

	b := box.NewBox()

	b.Use(MyCustomInterceptor)

	b.Handle("GET", "/articles/{article-id}", func(r *http.Request) *Article {
		return &Article{
			Id:    box.Param(r, "article-id"),
			Title: "Example",
		}
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Get(s.URL + "/articles/77")
	fmt.Println("Server:", resp.Header.Get("Server"))
	io.Copy(os.Stdout, resp.Body)

	// Output:
	// Server: MyServer
	// {"id":"77","title":"Example","text":"","created":"0001-01-01T00:00:00Z"}
}

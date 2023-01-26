package box_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/fulldump/box"
)

func ListArticles()  { /* ... */ }
func CreateArticle() { /* ... */ }
func GetArticle()    { /* ... */ }
func DeleteArticle() { /* ... */ }

var mockAccessLogPrintln = func(a ...interface{}) {
	a = append([]interface{}{"2023/01/26 21:05:15"}, a...)
	fmt.Println(a...)
}

func Example_Access_log_middleware() {

	// Mock access log
	box.DefaultAccessLogPrintln = mockAccessLogPrintln

	b := box.NewBox()

	b.Use(box.RecoverFromPanic) // use middlewares to print stacktraces

	b.Use(box.AccessLog)   // use middlewares to print logs
	b.Use(box.PrettyError) // use middlewares return pretty errors

	b.Handle("GET", "/articles", ListArticles)
	b.Handle("POST", "/articles", CreateArticle)
	b.Handle("GET", "/articles/{article-id}", GetArticle)
	b.Handle("DELETE", "/articles/{article-id}", DeleteArticle)

	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	http.Get(s.URL + "/articles")
	http.Get(s.URL + "/articles/25")
	http.Get(s.URL + "/articles/3")

	// Output:
	// 2023/01/26 21:05:15 GET /articles
	// 2023/01/26 21:05:15 GET /articles/25
	// 2023/01/26 21:05:15 GET /articles/3
}

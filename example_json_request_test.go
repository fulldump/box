package box_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/fulldump/box"
)

type CreateArticleRequest struct {
	Title string
	Text  string
}

type Article struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Text    string    `json:"text"`
	Created time.Time `json:"created"`
}

func Example_Json_request() {

	b := box.NewBox()
	b.Handle("POST", "/articles", func(input CreateArticleRequest) Article {
		fmt.Println("Persist new article...", input)
		return Article{
			Id:      "my-new-id",
			Title:   input.Title,
			Text:    input.Text,
			Created: time.Unix(1674762079, 0),
		}
	})
	s := httptest.NewServer(b)
	defer s.Close()
	// go b.ListenAndServe()

	resp, _ := http.Post(s.URL+"/articles", "application/json", strings.NewReader(`{
		"title": "My great article",
		"text": "Blah blah blah..."
	}`))
	io.Copy(os.Stdout, resp.Body)

	// Output:
	// Persist new article... {My great article Blah blah blah...}
	// {"id":"my-new-id","title":"My great article","text":"Blah blah blah...","created":"2023-01-26T20:41:19+01:00"}
}

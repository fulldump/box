# Box

<p align="center">
<a href="https://app.travis-ci.com/fulldump/box" rel="nofollow"><img src="https://app.travis-ci.com/fulldump/box.svg?branch=master" alt="Build Status"></a>
<a href="https://goreportcard.com/report/github.com/fulldump/box"><img src="https://goreportcard.com/badge/github.com/fulldump/box"></a>
<a href="https://godoc.org/github.com/fulldump/box"><img src="https://godoc.org/github.com/fulldump/box?status.svg" alt="GoDoc"></a>
</p>



Box is an HTTP router to speed up development. Box supports URL parameters, interceptors, magic handlers
and introspection documentation.

## Getting started

```go
package main

import (
	"github.com/fulldump/box"
)

func main() {
	
	b := box.NewBox()
	b.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World!"))
	})
	b.ListenAndServe() // listening at http://localhost:8080

}
```

## Sending JSON

```go
b := box.NewBox()

type MyResponse struct {
	Name string
	Age  int
}

b.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) MyResponse {
    return MyResponse{
        Name: "Fulanez",
        Age:  33,
    }
})
```

## URL parameters

```go
b := box.NewBox()

b.Handle("GET", "/articles/{article-id}", func(w http.ResponseWriter, r *http.Request) string {
    articleID := box.Param(r, "article-id")
    return "ArticleID is " + articleID
})
```

## Receiving and sending JSON

```go
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
```


package box

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBox2Http_ambiguousPathSemicolon(t *testing.T) {

	t.Run("Value with semicolon", func(t *testing.T) {

		b := NewBox()
		b.Handle("GET", "/resource/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := GetUrlParameter(r.Context(), "id")
			fmt.Fprint(w, id)
		})

		s := httptest.NewServer(b)

		res, err := http.Get(s.URL + "/resource/hello:world")
		AssertEqual(t, err, nil)
		AssertEqual(t, res.StatusCode, 200)
		body, _ := io.ReadAll(res.Body)
		AssertEqual(t, string(body), "hello:world")
	})

	t.Run("Operation with semicolon", func(t *testing.T) {

		b := NewBox()
		b.Resource("/resource/{id}").WithActions(
			Action(func(w http.ResponseWriter, r *http.Request) {
				id := GetUrlParameter(r.Context(), "id")
				fmt.Fprint(w, id)
			}).WithName("world"),
		)

		s := httptest.NewServer(b)

		res, err := http.Get(s.URL + "/resource/hello:world")
		AssertEqual(t, err, nil)
		AssertEqual(t, res.StatusCode, 200)
		body, _ := io.ReadAll(res.Body)
		AssertEqual(t, string(body), "hello")
	})

}

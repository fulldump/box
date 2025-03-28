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

func TestBox2Http_configurable404(t *testing.T) {

	b := NewBox()
	b.HandleResourceNotFound = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("HEY! Not Found"))
	}
	s := httptest.NewServer(b)

	res, _ := http.Get(s.URL + "/hello")
	AssertEqual(t, res.StatusCode, 404)
	body, _ := io.ReadAll(res.Body)
	AssertEqual(t, string(body), "HEY! Not Found")

}

func TestBox2Http_configurable405(t *testing.T) {

	b := NewBox()
	b.HandleMethodNotAllowed = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("HEY! Method Not Allowed"))
	}
	b.Handle("POST", "/hello", func(w http.ResponseWriter, r *http.Request) {})
	s := httptest.NewServer(b)

	res, _ := http.Get(s.URL + "/hello")
	AssertEqual(t, res.StatusCode, 405)
	body, _ := io.ReadAll(res.Body)
	AssertEqual(t, string(body), "HEY! Method Not Allowed")

}

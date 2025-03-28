package box

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestBox2Http_defaultSerializer(t *testing.T) {

	b := NewBox()
	b.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) any {
		return map[string]any{
			"hello": "world",
		}
	})
	s := httptest.NewServer(b)

	res, _ := http.Get(s.URL + "/hello")
	AssertEqual(t, res.StatusCode, http.StatusOK)
	AssertEqual(t, res.Header.Get("Content-Type"), "application/json")
	body, _ := io.ReadAll(res.Body)
	AssertEqual(t, string(body), "{\"hello\":\"world\"}\n")
}

func TestBox2Http_customSerializer(t *testing.T) {

	b := NewBox()
	b.Serializer = func(ctx context.Context, w io.Writer, v interface{}) error {

		resp := GetBoxContext(ctx).Response
		resp.Header().Set("Content-Type", "application/json+customwrap")
		resp.WriteHeader(http.StatusFound)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"wrap": v,
		})

		return nil
	}
	b.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) any {
		return map[string]any{
			"hello": "world",
		}
	})
	s := httptest.NewServer(b)

	res, _ := http.Get(s.URL + "/hello")
	AssertEqual(t, res.StatusCode, http.StatusFound)
	AssertEqual(t, res.Header.Get("Content-Type"), "application/json+customwrap")
	body, _ := io.ReadAll(res.Body)
	AssertEqual(t, string(body), "{\"wrap\":{\"hello\":\"world\"}}\n")
}

func TestBox2Http_customDeserializer(t *testing.T) {

	ErrPayloadTooLong := fmt.Errorf("ERROR! payload is too long")

	b := NewBox()
	b.WithInterceptors(PrettyError)
	b.Deserializer = func(ctx context.Context, r io.Reader, v interface{}) error {

		payload, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		if len(payload) > 15 {
			return ErrPayloadTooLong
		}

		return json.Unmarshal(payload, v)
	}
	type MyInput struct {
		Name string `json:"name"`
	}
	b.Handle("POST", "/hello", func(w http.ResponseWriter, r *http.Request, input *MyInput) any {
		return map[string]any{
			"hello": input.Name,
		}
	})
	s := httptest.NewServer(b)

	t.Run("Payload too long", func(t *testing.T) {
		res, _ := http.Post(s.URL+"/hello", "application/json", strings.NewReader(`too loooooooong string`))
		body, _ := io.ReadAll(res.Body)
		AssertEqual(t, string(body), ErrPayloadTooLong.Error())
	})

	t.Run("Payload OK", func(t *testing.T) {
		res, _ := http.Post(s.URL+"/hello", "application/json", strings.NewReader(`{"name":"U"}`))
		body, _ := io.ReadAll(res.Body)
		AssertEqual(t, string(body), "{\"hello\":\"U\"}\n")
	})

}

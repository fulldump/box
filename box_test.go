package box

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestNewBox_HappyPath(t *testing.T) {

	b := NewBox()

	b.Resource("/say-hello").
		WithActions(
			Get(func() string {
				return "Name World"
			}),
		)

	h := Box2Http(b)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/say-hello", nil)
	h.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)

	if "\"Name World\"\n" != string(body) {
		t.Error("Body does not match")
	}

}

func TestNewBox_PrintError(t *testing.T) {

	b := NewBox()

	b.WithInterceptors(PrettyError)

	b.Resource("/say-error").
		WithActions(
			Get(func() error {
				return errors.New("This is my error")
			}),
		)

	h := Box2Http(b)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/say-error", nil)
	h.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)

	if "This is my error" != string(body) {
		t.Error("Response does not match")
	}

}

func TestNewBox_EscapedUrlPath(t *testing.T) {

	b := NewBox()

	b.Resource("/{value}/hello").
		WithActions(
			Get(func(ctx context.Context) {
				value := GetUrlParameter(ctx, "value")
				_, _ = fmt.Fprintf(GetResponse(ctx), "Name '%s'", value)
			}),
		)

	h := Box2Http(b)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/Mr%20hello+world%2Fmoon/hello", nil)
	h.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)

	if "Name 'Mr hello world/moon'" != string(body) {
		t.Error("Body does not match")
	}

}

func TestNewBox_MethodAny(t *testing.T) {

	b := NewBox()

	b.Resource("/say-hello").
		WithActions(
			Post(func() string {
				return "Name World"
			}),
			AnyMethod(func() string {
				return "Any"
			}),
		)

	h := Box2Http(b)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/say-hello", nil)
	h.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)
	if `"Any"`+"\n" != string(body) {
		t.Error("Body does not match")
	}

}

func TestNewBox_ErrResourceNotFound(t *testing.T) {

	b := NewBox()

	b.WithInterceptors(
		func(next H) H {
			return func(ctx context.Context) {
				next(ctx)
				err := GetError(ctx)
				if err != ErrResourceNotFound {
					t.Error("Err should be ErrResourceNotFound")
				}
			}
		},
	)

	b.Resource("/hello").
		WithActions(Get(func() string {
			return "hello"
		}))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/invented", nil)
	b.ServeHTTP(w, r)

}

func TestNewBox_ErrMethodNotAllowed(t *testing.T) {

	b := NewBox()

	b.WithInterceptors(
		func(next H) H {
			return func(ctx context.Context) {
				next(ctx)
				err := GetError(ctx)
				if err != ErrMethodNotAllowed {
					t.Error("Err should be ErrMethodNotAllowed")
				}
			}
		},
	)

	b.Resource("/hello").
		WithActions(Get(func() string {
			return "hello"
		}))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/hello", nil)
	b.ServeHTTP(w, r)

}

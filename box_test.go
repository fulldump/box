package box

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestNewBox_HappyPath(t *testing.T) {

	b := NewBox()

	b.Resource("/say-hello").
		WithActions(
			Get(func() string {
				return "Hello World"
			}),
		)

	h := Box2Http(b)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/say-hello", nil)
	h.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)

	if "\"Hello World\"\n" != string(body) {
		t.Error("Body does not match")
	}

}

func TestNewBox_PrintError(t *testing.T) {

	b := NewBox()

	b.WithInterceptors(InterceptorPrintError)

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

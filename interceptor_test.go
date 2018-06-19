package box

import (
	"context"
	"net/http/httptest"
	"testing"
)

func TestInterceptors(t *testing.T) {

	b := NewBox()

	output := ""

	var wrap = func(label string) I {
		return func(next H) H {
			return func(ctx context.Context) {
				output += label + "["
				next(ctx)
				output += "]" + label
			}
		}
	}

	b.WithInterceptors(
		wrap("A"),
		wrap("B"),
	)

	b.Resource("/hello").
		WithInterceptors(
			wrap("C"),
			wrap("D"),
		).
		WithActions(
			Get(func() {
				output += "world"
			}).
				WithInterceptors(
					wrap("E"),
					wrap("F"),
				),
		)

	h := Box2Http(b)

	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("GET", "/hello", nil)

	h.ServeHTTP(w, r)

	if "A[B[C[D[E[F[world]F]E]D]C]B]A" != output {
		t.Error("Output do not match")
	}
}

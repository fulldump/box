//go:build !go1.22

package box

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBox2Http(t *testing.T) {

	t.Run("Path parameters", func(t *testing.T) {

		b := NewBox()
		b.Handle("GET", "/things/{thing_id}/history/{history_id}", func(ctx context.Context) {

			thingId := GetUrlParameter(ctx, "thing_id")
			AssertEqual(t, thingId, "THING-1")

			historyId := GetUrlParameter(ctx, "history_id")
			AssertEqual(t, historyId, "HISTORY-1")
		})

		s := httptest.NewServer(b)

		res, err := http.Get(s.URL + "/things/THING-1/history/HISTORY-1")
		AssertEqual(t, err, nil)
		AssertEqual(t, res.StatusCode, 200)
	})

}

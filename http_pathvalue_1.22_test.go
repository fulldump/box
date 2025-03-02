//go:build go1.22

package box

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBox2Http_pathValue(t *testing.T) {

	t.Run("Path parameters", func(t *testing.T) {

		b := NewBox()
		b.Handle("GET", "/things/{thing_id}/history/{history_id}", func(w http.ResponseWriter, r *http.Request) {

			thingId := r.PathValue("thing_id")
			AssertEqual(t, thingId, "THING-1")

			historyId := r.PathValue("history_id")
			AssertEqual(t, historyId, "HISTORY-1")
		})

		s := httptest.NewServer(b)

		res, err := http.Get(s.URL + "/things/THING-1/history/HISTORY-1")
		AssertEqual(t, err, nil)
		AssertEqual(t, res.StatusCode, 200)
	})

}

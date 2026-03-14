package box_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fulldump/box"
	"github.com/gorilla/websocket"
)

func TestHandleWebSocketEcho(t *testing.T) {
	b := box.NewBox()
	b.HandleWebSocket("/ws", func(_ context.Context, conn *websocket.Conn) {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			return
		}
		_ = conn.WriteMessage(messageType, payload)
	})

	s := httptest.NewServer(b)
	defer s.Close()

	wsURL := "ws" + strings.TrimPrefix(s.URL, "http") + "/ws"

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("dial websocket: %v", err)
	}
	defer conn.Close()

	const expected = "ping"
	if err := conn.WriteMessage(websocket.TextMessage, []byte(expected)); err != nil {
		t.Fatalf("write websocket: %v", err)
	}

	_, got, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("read websocket: %v", err)
	}

	if string(got) != expected {
		t.Fatalf("unexpected message: got %q expected %q", string(got), expected)
	}
}

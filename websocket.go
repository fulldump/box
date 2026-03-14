package box

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler func(ctx context.Context, conn *websocket.Conn)

type WebSocketErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

type WebSocketOptions struct {
	Upgrader websocket.Upgrader
	OnError  WebSocketErrorHandler
}

func DefaultWebSocketOptions() WebSocketOptions {
	return WebSocketOptions{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		OnError: func(w http.ResponseWriter, _ *http.Request, err error) {
			http.Error(w, fmt.Sprintf("websocket upgrade failed: %v", err), http.StatusBadRequest)
		},
	}
}

func (r *R) HandleWebSocket(path string, handler WebSocketHandler, options ...WebSocketOptions) *A {
	wsOptions := DefaultWebSocketOptions()
	if len(options) > 0 {
		wsOptions = options[0]
		if wsOptions.OnError == nil {
			wsOptions.OnError = DefaultWebSocketOptions().OnError
		}
	}

	return r.Handle(http.MethodGet, path, func(w http.ResponseWriter, req *http.Request) {
		conn, err := wsOptions.Upgrader.Upgrade(w, req, nil)
		if err != nil {
			wsOptions.OnError(w, req, err)
			return
		}
		defer conn.Close()

		handler(req.Context(), conn)
	})
}

# WebSockets

`box` includes a dedicated helper to register WebSocket routes.

## Echo server example

```go
import (
	"context"

	"github.com/gorilla/websocket"
)

b := box.NewBox()

b.HandleWebSocket("/ws/echo", func(_ context.Context, conn *websocket.Conn) {
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			return
		}
		_ = conn.WriteMessage(messageType, payload)
	}
})
```

## Custom upgrader

Use `box.WebSocketOptions` to customize buffers, origin checks, and errors.

```go
opts := box.WebSocketOptions{
	Upgrader: websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return strings.HasSuffix(r.Host, ".my-domain.com")
		},
	},
	OnError: func(w http.ResponseWriter, _ *http.Request, err error) {
		http.Error(w, "websocket upgrade refused", http.StatusForbidden)
		log.Printf("ws upgrade error: %v", err)
	},
}

b.HandleWebSocket("/ws", ChatHandler, opts)
```

## Production advice

- Enforce a strict `CheckOrigin` policy.
- Set read/write deadlines in handlers.
- Handle ping/pong and close frames explicitly.
- Backpressure outbound writes to avoid memory spikes.

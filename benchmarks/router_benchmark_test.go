package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fulldump/box"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
)

type benchmarkCase struct {
	handler http.Handler
	req     func() *http.Request
}

type createItemRequest struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type createItemResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

func BenchmarkRouters(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	b.Run("StaticGET", func(b *testing.B) {
		runRouters(b, map[string]benchmarkCase{
			"box":  boxStaticCase(),
			"gin":  ginStaticCase(),
			"chi":  chiStaticCase(),
			"echo": echoStaticCase(),
		})
	})

	b.Run("PathParamGET", func(b *testing.B) {
		runRouters(b, map[string]benchmarkCase{
			"box":  boxPathParamCase(),
			"gin":  ginPathParamCase(),
			"chi":  chiPathParamCase(),
			"echo": echoPathParamCase(),
		})
	})

	b.Run("JSONPOST", func(b *testing.B) {
		runRouters(b, map[string]benchmarkCase{
			"box":  boxJSONCase(),
			"gin":  ginJSONCase(),
			"chi":  chiJSONCase(),
			"echo": echoJSONCase(),
		})
	})
}

func runRouters(b *testing.B, routerCases map[string]benchmarkCase) {
	for routerName, testCase := range routerCases {
		routerName := routerName
		testCase := testCase

		b.Run(routerName, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				w := httptest.NewRecorder()
				testCase.handler.ServeHTTP(w, testCase.req())
				if w.Code != http.StatusOK {
					b.Fatalf("unexpected status code %d", w.Code)
				}
			}
		})
	}
}

func boxStaticCase() benchmarkCase {
	b := box.NewBox()
	b.HandleFunc(http.MethodGet, "/hello", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	return benchmarkCase{
		handler: b,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/hello", nil)
		},
	}
}

func ginStaticCase() benchmarkCase {
	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return benchmarkCase{
		handler: r,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/hello", nil)
		},
	}
}

func chiStaticCase() benchmarkCase {
	r := chi.NewRouter()
	r.Get("/hello", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	return benchmarkCase{
		handler: r,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/hello", nil)
		},
	}
}

func echoStaticCase() benchmarkCase {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	return benchmarkCase{
		handler: e,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/hello", nil)
		},
	}
}

func boxPathParamCase() benchmarkCase {
	b := box.NewBox()
	b.HandleFunc(http.MethodGet, "/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(r.PathValue("id")))
	})

	return benchmarkCase{
		handler: b,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/users/42", nil)
		},
	}
}

func ginPathParamCase() benchmarkCase {
	r := gin.New()
	r.GET("/users/:id", func(c *gin.Context) {
		c.String(http.StatusOK, c.Param("id"))
	})

	return benchmarkCase{
		handler: r,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/users/42", nil)
		},
	}
}

func chiPathParamCase() benchmarkCase {
	r := chi.NewRouter()
	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(chi.URLParam(r, "id")))
	})

	return benchmarkCase{
		handler: r,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/users/42", nil)
		},
	}
}

func echoPathParamCase() benchmarkCase {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})

	return benchmarkCase{
		handler: e,
		req: func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/users/42", nil)
		},
	}
}

func boxJSONCase() benchmarkCase {
	b := box.NewBox()
	b.Handle(http.MethodPost, "/items", func(in createItemRequest) createItemResponse {
		return createItemResponse{
			ID:   "i-1",
			Name: in.Name,
			Qty:  in.Qty + 1,
		}
	})

	payload, _ := json.Marshal(createItemRequest{Name: "apple", Qty: 1})

	return benchmarkCase{
		handler: b,
		req: func() *http.Request {
			req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			return req
		},
	}
}

func ginJSONCase() benchmarkCase {
	r := gin.New()
	r.POST("/items", func(c *gin.Context) {
		var in createItemRequest
		if err := c.BindJSON(&in); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, createItemResponse{
			ID:   "i-1",
			Name: in.Name,
			Qty:  in.Qty + 1,
		})
	})

	payload, _ := json.Marshal(createItemRequest{Name: "apple", Qty: 1})

	return benchmarkCase{
		handler: r,
		req: func() *http.Request {
			req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			return req
		},
	}
}

func chiJSONCase() benchmarkCase {
	r := chi.NewRouter()
	r.Post("/items", func(w http.ResponseWriter, req *http.Request) {
		var in createItemRequest
		if err := json.NewDecoder(req.Body).Decode(&in); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(createItemResponse{
			ID:   "i-1",
			Name: in.Name,
			Qty:  in.Qty + 1,
		})
	})

	payload, _ := json.Marshal(createItemRequest{Name: "apple", Qty: 1})

	return benchmarkCase{
		handler: r,
		req: func() *http.Request {
			req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			return req
		},
	}
}

func echoJSONCase() benchmarkCase {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.POST("/items", func(c echo.Context) error {
		var in createItemRequest
		if err := c.Bind(&in); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.JSON(http.StatusOK, createItemResponse{
			ID:   "i-1",
			Name: in.Name,
			Qty:  in.Qty + 1,
		})
	})

	payload, _ := json.Marshal(createItemRequest{Name: "apple", Qty: 1})

	return benchmarkCase{
		handler: e,
		req: func() *http.Request {
			req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			return req
		},
	}
}

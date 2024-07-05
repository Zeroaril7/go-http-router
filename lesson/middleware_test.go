package lesson

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	m.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {

	expect := "Middleware"

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := LogMiddleware{Handler: router}

	req := httptest.NewRequest(http.MethodGet, host, nil)
	rec := httptest.NewRecorder()

	middleware.ServeHTTP(rec, req)
	res := rec.Result()

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, expect, string(bytes))
}

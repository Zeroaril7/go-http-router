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

const (
	pathProduct = "product/"
)

func TestParams(t *testing.T) {

	expect := "Product 1"

	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		txt := "Product " + id
		fmt.Fprint(w, txt)
	})

	req := httptest.NewRequest(http.MethodGet, host+pathProduct+"1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, expect, string(bytes))
}

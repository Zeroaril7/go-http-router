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
	pathItem  = "/item/"
	pathImage = "images/"
)

func TestNamedParams(t *testing.T) {

	expect := "Product 1 item 1"

	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		txt := "Product " + id + " item " + itemId
		fmt.Fprint(w, txt)
	})

	req := httptest.NewRequest(http.MethodGet, host+pathProduct+"1"+pathItem+"1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, expect, string(bytes))
}

func TestCatchAllParams(t *testing.T) {

	expect := "Image: /test.png"

	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		txt := "Image: " + image
		fmt.Fprint(w, txt)
	})

	req := httptest.NewRequest(http.MethodGet, host+pathImage+"test.png", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, expect, string(bytes))
}

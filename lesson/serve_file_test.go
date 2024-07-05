package lesson

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

const (
	pathFile = "files/"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {

	expect := "test"

	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	req := httptest.NewRequest(http.MethodGet, host+pathFile+"test.txt", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, expect, string(bytes))
}

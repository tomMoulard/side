package function

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed html
	html embed.FS

	handler http.Handler
)

func init() {
	fsys, err := fs.Sub(html, "html")
	if err != nil {
		panic(err)
	}

	handler = http.FileServer(http.FS(fsys))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}

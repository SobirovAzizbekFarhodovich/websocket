package transport

import (
	"net/http"
	"path/filepath"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("web", "index.html"))
}

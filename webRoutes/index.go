package webRoutes

import (
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Read the index.html
	data, _ := os.ReadFile("index.html")

	// Write it to stdout
	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write(data)
}

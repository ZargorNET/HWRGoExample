package webRoutes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloName(w http.ResponseWriter, r *http.Request) {
	// Get the variables out of the url defined in the main
	vars := mux.Vars(r)

	_, _ = fmt.Fprintf(w, "Hello \"%s\"", vars["name"])
}

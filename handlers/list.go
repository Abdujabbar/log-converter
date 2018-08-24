package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func list(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("CurrentPage: %v", params.ByName("page"))))
}

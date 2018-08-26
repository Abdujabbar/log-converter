package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var tableTmpl = "<table border=1 cellpadding=10 cellspacing=5 style='width: 100%%'><thead><th>ID</th><th>TIME</th><th>MSG</th><th>FORMAT</th></thead><tbody>%s</tbody><tfoot>%s</tfoot></table>"
var rawTmpl = "<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>"
var tfootTpml = "<tr><td colspan=3>Total Rows</td><td>%v</td></tr>"

func list(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	page := params.ByName("page")
	currentPage := 1
	if v, err := strconv.Atoi(page); err == nil {
		currentPage = v
	}
	limit := 50
	offset := (currentPage - 1) * limit
	raws, err := provider.FindAll(limit, offset)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while fetching rows: %v", err.Error())))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tableBody := ""
	for _, v := range raws {
		t := time.Unix(v.Time, 0)
		tableBody += fmt.Sprintf(rawTmpl, v.ID, t.Format(time.RFC1123), v.Msg, v.Format)
	}
	total, _ := provider.Count()
	tfoot := fmt.Sprintf(tfootTpml, total)
	table := fmt.Sprintf(tableTmpl, tableBody, tfoot)
	w.Write([]byte(table))
}

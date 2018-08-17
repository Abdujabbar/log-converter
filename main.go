package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/Abdujabbor/log-converter/repository"
)

var tableTmpl = "<table border=1 cellpadding=10 cellspacing=5 style='width: 100%%'><thead><th>ID</th><th>TIME</th><th>MSG</th><th>FORMAT</th></thead><tbody>%s</tbody><tfoot>%s</tfoot></table>"
var rawTmpl = "<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>"
var tfootTpml = "<tr><td colspan=3>Total Rows</td><td>%v</td></tr>"
var dao repository.DAO

func main() {
	fmt.Println("Please wait until servers will be ready")
	dao = repository.DAO{
		Server:   "localhost",
		Database: "plogs",
	}

	err := dao.Connect()
	dao.Truncate()
	if err != nil {
		panic(err)
	}

	go startMonitoringFiles(&dao, os.Args)

	router := httprouter.New()

	router.GET("/logs", logsHandler)

	router.GET("/logs/:page", logsHandler)

	http.ListenAndServe(":8080", router)
}

func logsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	page := params.ByName("page")
	currentPage := 1
	if v, err := strconv.Atoi(page); err == nil {
		currentPage = v
	}
	limit := 10
	offset := (currentPage - 1) * limit
	raws, err := dao.FindAll(limit, offset)
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
	total, _ := dao.Count()
	tfoot := fmt.Sprintf(tfootTpml, total)
	table := fmt.Sprintf(tableTmpl, tableBody, tfoot)
	w.Write([]byte(table))
}

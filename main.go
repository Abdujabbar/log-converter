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

const ftype = "1"
const stype = "2"

var dao = repository.DAO{
	Server:   "localhost",
	Database: "testdb",
}
var tableTmpl = "<table border=1 cellpadding=10 cellspacing=5 style='width: 100%%'><thead><th>ID</th><th>TIME</th><th>MSG</th><th>FORMAT</th></thead><tbody>%s</tbody></table>"
var rawTmpl = "<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>"

func serveLogsList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
	table := fmt.Sprintf(tableTmpl, tableBody)
	w.Write([]byte(table))
}

func main() {
	fmt.Println("Please wait until servers will be ready")
	err := dao.Connect()
	if err != nil {
		panic(err)
	}

	go process(os.Args)

	router := httprouter.New()

	router.GET("/list/:page", serveLogsList)

	http.ListenAndServe(":8080", router)
}

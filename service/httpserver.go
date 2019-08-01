package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	configStruct "github.com/vsel/goSqlWeb/config/struct"
)

// ListenHTTP create http server with routing
func ListenHTTP(config configStruct.Configuration) error {
	r := mux.NewRouter()
	r.HandleFunc("/test", test)
	http.Handle("/", r)
	fmt.Println(`Starting up on :{{.config.HTTPServer.Port}}`)
	err := http.ListenAndServe(":"+config.HTTPServer.Port, nil)
	return err
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Test")
}

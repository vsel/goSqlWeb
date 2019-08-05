package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	configStruct "github.com/vsel/goSqlWeb/config/struct"
)

// ListenHTTP create http server with routing
func ListenHTTP(config configStruct.Configuration, env *Env) error {
	r := mux.NewRouter()
	r.HandleFunc("/test", env.test)
	http.Handle("/", r)
	fmt.Println("Starting up on :", config.HTTPServer.Port)
	err := http.ListenAndServe(":"+config.HTTPServer.Port, nil)
	return err
}

func (env *Env) test (w http.ResponseWriter, req *http.Request) {
	res, err := env.DB.GetTestData()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
	fmt.Fprintln(w, res)
}

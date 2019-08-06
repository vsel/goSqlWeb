package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"

	_ "github.com/doug-martin/goqu/v8/dialect/postgres"
	service "github.com/vsel/goSqlWeb/service"
)

func main() {
	config := service.GetConfig(".")

	db, err := service.ConnectToDB(config)
	if err != nil {
		log.Fatal(err)
	}

	env := &service.Env{db}
	fmt.Println(env.DB.GetTestData())

	err = service.ListenHTTP(config, env)
	if err != nil {
		log.Fatal(err)
	}

}

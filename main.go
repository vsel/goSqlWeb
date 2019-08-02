package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"

	"github.com/doug-martin/goqu/v8"
	_ "github.com/doug-martin/goqu/v8/dialect/postgres"
	service "github.com/vsel/goSqlWeb/service"
)

func main() {
	config := service.GetConfig(".")

	db, err := service.ConnectToDB(config)
	if err != nil {
		log.Fatal(err)
	}

	err = service.ListenHTTP(config)
	if err != nil {
		log.Fatal(err)
	}

	dialect := goqu.Dialect("postgres")

	comments, err := service.GetTestData(dialect, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(comments)

}

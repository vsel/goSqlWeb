package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/doug-martin/goqu/v8"
	_ "github.com/doug-martin/goqu/v8/dialect/postgres"
	source "github.com/vsel/goSqlWeb/source"
)

func main() {
	config := source.GetConfig(".")

	db, err := source.ConnectToDB(config)
	if err != nil {
		log.Fatal(err)
	}

	dialect := goqu.Dialect("postgres")

	err = source.InitTables(dialect, db)
	if err != nil {
		log.Fatal(err)
	}
}

package source

import (
	"bytes"
	"database/sql"
	"fmt"
	"text/template"

	"github.com/doug-martin/goqu/v8"

	configStruct "github.com/vsel/goSqlWeb/config/struct"
)

func getConnectionString(config configStruct.Configuration) (string, error) {
	input := configStruct.DatabaseConfiguration{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		DBName:   config.Database.DBName,
		SSL:      config.Database.SSL,
	}

	connectionString := `postgresql://{{.User}}:{{.Password}}@{{.Host}}:{{.Port}}/{{.DBName}}?sslmode={{.SSL}}`
	templateString := template.Must(template.New("connectionString").Parse(connectionString))
	var stringParsed bytes.Buffer
	if errExecute := templateString.Execute(&stringParsed, input); errExecute != nil {
		return "", errExecute
	}

	return stringParsed.String(), nil
}

// ConnectToDB create db connection
func ConnectToDB(config configStruct.Configuration) (*sql.DB, error) {
	connStr, err := getConnectionString(config)
	if err != nil {
		return nil, err
	}

	fmt.Println("connecting to: " + connStr)
	return sql.Open("postgres", connStr)
}

// InitTables init tables
func InitTables(dialect goqu.DialectWrapper, db *sql.DB) error {
	createRows, err := db.Query(`
	CREATE TABLE IF NOT EXISTS public.comments_test
	(
		id bigserial NOT NULL,
		data text,
		author_id bigint,
		PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	);

	ALTER TABLE public.comments_test
		OWNER to "postgres";
	`)
	if err != nil {
		return err
	}
	defer createRows.Close()
	return nil
}

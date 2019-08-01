package source

import (
	"bytes"
	"database/sql"
	"fmt"
	"text/template"

	"github.com/doug-martin/goqu/v8"

	configStruct "github.com/vsel/goSqlWeb/config/struct"
	models "github.com/vsel/goSqlWeb/models"
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

// GetTestData get test data
func GetTestData(dialect goqu.DialectWrapper, db *sql.DB) ([]models.Comments, error) {
	dialectString := dialect.From("comments").Where(goqu.Ex{"author_id": 1})
	query, args, err := dialectString.ToSQL()
	if err != nil {
		fmt.Println("Failed to generate query string", err.Error())
	} else {
		fmt.Println(query, args)
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	comments := []models.Comments{}
	for rows.Next() {
		var c models.Comments
		err = rows.Scan(&c.ID, &c.Data, &c.AuthorID)
		if err != nil {
			fmt.Println("Failed to get row", err.Error())
		}
		comments = append(comments, c)
	}

	return comments, nil
}

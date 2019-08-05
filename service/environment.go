package service
import (
	models "github.com/vsel/goSqlWeb/models"
)

// Env is
type Env struct {
	DB Datastore
}

// Datastore is interface for access to db
type Datastore interface {
	GetTestData() ([]models.Comments, error)
}

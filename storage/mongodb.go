package storage

import "github.com/Devon-ODell/testcode/models"

type MongoStorage struct{}

func (s *MongoStorage) Get(id int) *models.User {
	return &models.User{
		ID:   1,
		Name: "Foo",
	}
}

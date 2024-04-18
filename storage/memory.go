package storage

import "github.com/Devon-ODell/testcode/models"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *models.User {
	return &models.User{
		ID:   1,
		Name: "Foo",
	}
}

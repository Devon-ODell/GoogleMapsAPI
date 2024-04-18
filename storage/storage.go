package storage

import "github.com/Devon-ODell/testcode/models"

type Storage interface {
	Get(int) *models.User
}




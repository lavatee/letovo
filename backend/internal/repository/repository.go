package repository

import backend "github.com/lavatee/children_backend"

type Repository struct {
	Children Children
}

type Children interface {
	TakeChild(childId int, userEmail string, userFirstName string, userLastName string, userPhone string, userTelegram string, userClass string) (string, string, error)
	GetChildren() []backend.Child
}

func NewRepository(db map[int]backend.Child) *Repository {
	return &Repository{
		Children: NewChildrenMap(db),
	}
}

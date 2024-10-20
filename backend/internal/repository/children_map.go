package repository

import (
	"errors"
	"fmt"

	backend "github.com/lavatee/children_backend"
)

type ChildrenMap struct {
	db map[int]backend.Child
}

func NewChildrenMap(db map[int]backend.Child) *ChildrenMap {
	return &ChildrenMap{
		db: db,
	}
}

func (r *ChildrenMap) TakeChild(childId int, userEmail string, userFirstName string, userLastName string, userPhone string, userTelegram string, userClass string) (string, string, error) {
	for _, child := range r.db {
		if child.UserEmail == userEmail {
			return "", "", errors.New("user has already taken a child")
		}
	}
	if _, ok := r.db[childId]; !ok {
		return "", "", fmt.Errorf("child with id %d does not in map", childId)
	}
	if r.db[childId].IsTaken {
		return "", "", fmt.Errorf("child with id %d has already been taken", childId)
	}
	child := r.db[childId]
	child.IsTaken = true
	child.UserEmail = userEmail
	child.UserFirstName = userFirstName
	child.UserLastName = userLastName
	child.UserClass = userClass
	child.UserTelegram = userTelegram
	child.UserPhoneNumber = userPhone
	r.db[childId] = child
	return child.FirstName + " " + child.LastName, child.Gift, nil
}

func (r *ChildrenMap) GetChildren() []backend.Child {
	children := make([]backend.Child, 0)
	for _, child := range r.db {
		children = append(children, child)
	}
	return children
}

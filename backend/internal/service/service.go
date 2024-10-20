package service

import (
	"net/smtp"

	backend "github.com/lavatee/children_backend"
	"github.com/lavatee/children_backend/internal/repository"
)

type Service struct {
	Children Children
}

type Children interface {
	TakeChild(childId int, userEmail string, userFirstName string, userLastName string, userPhone string, userTelegram string, userClass string, code string) (string, string, error)
	GetChildren() []backend.Child
	SendCode(userEmail string) error
	NewAdmin(email string) error
	GetChildrenInfo(email string, code string) ([]backend.Child, error)
}

func NewService(repo *repository.Repository, auth smtp.Auth, gmail string, host string, port string) *Service {
	return &Service{
		Children: NewChildrenService(repo, auth, gmail, host, port),
	}
}

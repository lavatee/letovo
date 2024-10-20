package service

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"sync"
	"time"

	backend "github.com/lavatee/children_backend"
	"github.com/lavatee/children_backend/internal/repository"
)

type CodeStore struct {
	mu    sync.RWMutex
	codes map[string][2]interface{}
}

func NewCodeStore() *CodeStore {
	return &CodeStore{
		codes: make(map[string][2]interface{}),
	}
}

func (cs *CodeStore) SetCode(userEmail string, code string) {
	expiration := time.Now().Add(60 * time.Second)
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.codes[userEmail] = [2]interface{}{code, expiration}
}

func (cs *CodeStore) VerifyCode(userEmail string, code string) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	data, exists := cs.codes[userEmail]
	if !exists {
		return false
	}

	storedCode := data[0].(string)
	expiration := data[1].(time.Time)
	if time.Now().After(expiration) {
		delete(cs.codes, userEmail)
		return false
	}
	return storedCode == code
}

type ChildrenService struct {
	Repo      *repository.Repository
	SmtpAuth  smtp.Auth
	Gmail     string
	SmtpHost  string
	SmtpPort  string
	CodeStore *CodeStore
	Admins    []string
}

func NewChildrenService(repo *repository.Repository, auth smtp.Auth, gmail string, host string, port string) *ChildrenService {
	return &ChildrenService{
		Repo:      repo,
		SmtpAuth:  auth,
		Gmail:     gmail,
		SmtpHost:  host,
		SmtpPort:  port,
		CodeStore: NewCodeStore(),
		Admins:    []string{"aleksgraznov0@gmail.com", "julis.009@mail.ru"},
	}
}

func generateRandomCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}

func (s *ChildrenService) SendCode(userEmail string) error {
	code := strconv.Itoa(generateRandomCode())
	if err := smtp.SendMail(s.SmtpHost+":"+s.SmtpPort, s.SmtpAuth, s.Gmail, []string{userEmail}, []byte(fmt.Sprintf("Subject: %s", code))); err != nil {
		return err
	}
	s.CodeStore.SetCode(userEmail, code)
	return nil
}

func (s *ChildrenService) TakeChild(childId int, userEmail string, userFirstName string, userLastName string, userPhone string, userTelegram string, userClass string, code string) (string, string, error) {
	if !s.CodeStore.VerifyCode(userEmail, code) {
		return "", "", fmt.Errorf("invalid code: %s", code)
	}
	fullName, gift, err := s.Repo.Children.TakeChild(childId, userEmail, userFirstName, userLastName, userPhone, userTelegram, userClass)
	if err != nil {
		return "", "", err
	}
	err = s.SendMessage(userEmail, fullName, gift)
	return fullName, gift, err
}

func (s *ChildrenService) SendMessage(email string, childFullname string, gift string) error {
	if err := smtp.SendMail(s.SmtpHost+":"+s.SmtpPort, s.SmtpAuth, s.Gmail, []string{email}, []byte(fmt.Sprintf("Subject: Вы дарите %s ребенку %s", gift, childFullname))); err != nil {
		return err
	}
	return nil
}

func (s *ChildrenService) GetChildren() []backend.Child {
	return s.Repo.Children.GetChildren()
}

func (s *ChildrenService) NewAdmin(email string) error {
	isAlreadyAdmin := false
	for _, admin := range s.Admins {
		if admin == email {
			isAlreadyAdmin = true
		}
	}
	if isAlreadyAdmin {
		return fmt.Errorf("email %s already in list of admins", email)
	}
	admins := s.Admins
	admins = append(admins, email)
	s.Admins = admins
	return nil
}

func (s *ChildrenService) GetChildrenInfo(email string, code string) ([]backend.Child, error) {
	isAdmin := false
	for _, admin := range s.Admins {
		if admin == email {
			isAdmin = true
		}
	}
	if !isAdmin {
		return nil, fmt.Errorf("email %s not in list of admins", email)
	}
	if !s.CodeStore.VerifyCode(email, code) {
		return nil, fmt.Errorf("invalid code: %s", code)
	}
	children := s.Repo.Children.GetChildren()
	takenChildren := make([]backend.Child, 0)
	for _, child := range children {
		if child.IsTaken {
			takenChildren = append(takenChildren, child)
		}
	}
	return takenChildren, nil
}

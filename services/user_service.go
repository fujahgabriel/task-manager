package services

import (
	"errors"
	"sync"
	"task-manager/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users  []models.User
	mutex  sync.Mutex
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  []models.User{},
		nextID: 1,
	}
}

func (s *UserService) Register(email, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if email already exists
	for _, user := range s.users {
		if user.Email == email {
			return nil, errors.New("email already exists")
		}
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:       s.nextID,
		Email:    email,
		Password: string(hashedPassword),
	}
	s.users = append(s.users, user)
	s.nextID++
	return &user, nil
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Email == email {
			// Compare the provided password with the hashed password
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return nil, errors.New("incorrect password")
			}
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

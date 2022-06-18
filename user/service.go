package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	//TODO implement me
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	user := User{
		Name:         input.Name,
		Occupation:   input.Occupation,
		Email:        input.Email,
		PasswordHash: string(password),
		Role:         "user",
	}
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

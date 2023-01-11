package service

import (
	"context"
	"github.com/bl4ckf1sher/ad-service/internal/domain"
	"github.com/bl4ckf1sher/ad-service/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type User struct {
	userRepo repositories.User
}

func NewUsersService(repo repositories.User) *User {
	return &User{repo}
}

func (s User) GetUsers(c context.Context) (*[]domain.User, error) {
	var users *[]domain.User

	users, err := s.userRepo.GetAll(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s User) GetUserById(c context.Context, id uuid.UUID) (*domain.User, error) {
	var user *domain.User

	user, err := s.userRepo.Get(c, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s User) CreateUser(c context.Context, user domain.User) (err error) {
	err = s.userRepo.Create(c, user)
	return
}

func (s User) DeleteUser(c context.Context, id uuid.UUID) (err error) {
	err = s.userRepo.Delete(c, id)
	if err != nil {
		return
	}

	return nil
}

func (s User) UpdateUser(c context.Context, user domain.User) (err error) {
	err = s.userRepo.Update(c, user)
	return
}

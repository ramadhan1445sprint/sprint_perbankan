package svc

import (
	"github.com/ramadhan1445sprint/sprint_segokuning/crypto"
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
	"github.com/ramadhan1445sprint/sprint_segokuning/repo"
)

type UserSvc interface {
	Register(newUser entity.RegistrationPayload) (string, error)
	Login(user entity.Credential) (string, string, error)
}

type userSvc struct {
	repo repo.UserRepo
}

func NewUserSvc(repo repo.UserRepo) UserSvc {
	return &userSvc{repo}
}

func (s *userSvc) Register(newUser entity.RegistrationPayload) (string, error) {
	user := entity.NewUser(newUser.Email, newUser.Name, newUser.Password)

	if err := user.Validate(); err != nil {
		return "", customErr.NewBadRequestError(err.Error())
	}

	existingUser, err := s.repo.GetUser(newUser.Email)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return "", err
		}
	}

	if existingUser != nil {
		return "", customErr.NewConflictError("user already exist")
	}

	hashedPassword, err := crypto.GenerateHashedPassword(newUser.Password)
	if err != nil {
		return "", err
	}

	id, err := s.repo.CreateUser(&newUser, hashedPassword)
	if err != nil {
		return "", err
	}

	token, err := crypto.GenerateToken(id, newUser.Email, newUser.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userSvc) Login(creds entity.Credential) (string, string, error) {
	if err := creds.Validate(); err != nil {
		return "", "", customErr.NewBadRequestError(err.Error())
	}

	user, err := s.repo.GetUser(creds.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", "", customErr.NewNotFoundError("user not found")
		}
		return "", "", err
	}

	err = crypto.VerifyPassword(creds.Password, user.Password)
	if err != nil {
		return "", "", customErr.NewBadRequestError("wrong password!")
	}

	token, err := crypto.GenerateToken(user.Id, user.Email, user.Name)
	if err != nil {
		return "", "", customErr.NewBadRequestError(err.Error())
	}

	return user.Name, token, nil
}

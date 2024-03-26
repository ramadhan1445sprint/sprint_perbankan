package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
)

type UserRepo interface {
	GetUser(email string) (*entity.User, error)
	CreateUser(user *entity.RegistrationPayload, hashPassword string) (string, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) GetUser(email string) (*entity.User, error) {
	var user entity.User
	query := "SELECT * FROM users WHERE email = $1"

	err := r.db.Get(&user, query, email)
	if err != nil {
		fmt.Println("1", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CreateUser(user *entity.RegistrationPayload, hashPassword string) (string, error) {
	var id string

	statement := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRowx(statement, user.Name, user.Email, hashPassword)

	if err := row.Scan(&id); err != nil {
		fmt.Println("2", err)
		return "", err
	}

	return id, nil
}

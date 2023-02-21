package repositories

import (
	"fmt"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user *models.User) error
	Get() ([]models.User, error)
}

type users struct {
	DB *sqlx.DB
}

func NewUsers(db *sqlx.DB) Users {
	return &users{
		DB: db,
	}
}

func (u *users) Create(user *models.User) error {
	_, err := u.DB.NamedExec("INSERT INTO public.users(name, email) VALUES (:name, :email)", user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (u *users) Get() ([]models.User, error) {
	var usersSlice []models.User
	err := u.DB.Select(&usersSlice, "SELECT id, email, name FROM public.users")
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	return usersSlice, nil
}

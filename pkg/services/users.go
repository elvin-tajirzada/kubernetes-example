package services

import (
	"github.com/elvin-tacirzade/kubernetes-example/pkg/models"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/repositories"
	"log"
	"net/http"
)

type Users interface {
	Create(name, email string) *models.Response
	Get() *models.Response
}

type users struct {
	usersRepository repositories.Users
}

func NewUsers(usersRepository repositories.Users) Users {
	return &users{
		usersRepository: usersRepository,
	}
}

func (u *users) Create(name, email string) *models.Response {
	if name == "" || email == "" {
		log.Println("failed to create a new user: name or email fields can't be empty")
		return models.NewResponse(http.StatusBadRequest, map[string]interface{}{
			"status":  models.ResponseStatusFailed,
			"message": "name or email fields can't be empty",
		})
	}
	user := models.User{Name: name, Email: email}
	err := u.usersRepository.Create(&user)
	if err != nil {
		log.Println(err)
		return models.NewResponse(http.StatusInternalServerError, models.ResponseStatusFailedMap)
	}
	return models.NewResponse(http.StatusCreated, map[string]interface{}{
		"status": models.ResponseStatusSuccess,
	})
}

func (u *users) Get() *models.Response {
	getUsers, err := u.usersRepository.Get()
	if err != nil {
		log.Println(err)
		return models.NewResponse(http.StatusInternalServerError, models.ResponseStatusFailedMap)
	}
	return models.NewResponse(http.StatusOK, map[string]interface{}{
		"data": getUsers,
	})
}

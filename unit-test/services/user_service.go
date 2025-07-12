package services

import (
	"errors"
	"golang-journeys/unit-test/entities"
	"golang-journeys/unit-test/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (userService UserService) Get(id string) (*entities.User, error) {
	entity := userService.Repository.FindById(id)

	if entity != nil {
		return entity, nil
	}

	return nil, errors.New("entity is not found")
}

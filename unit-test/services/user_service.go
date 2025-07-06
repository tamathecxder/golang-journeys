package services

import (
	"errors"
	"golang-journeys/unit-test/entities"
	"golang-journeys/unit-test/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func (userService UserService) Find(id string) (*entities.User, error) {
	entity := userService.repository.FindById(id)

	if entity != nil {
		return entity, nil
	}

	return nil, errors.New("entity is not found")
}

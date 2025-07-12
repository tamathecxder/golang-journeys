package repositories

import (
	"golang-journeys/unit-test/entities"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindById(id string) *entities.User {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entities.User)
	return &category
}

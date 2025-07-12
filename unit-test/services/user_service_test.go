package services

import (
	"golang-journeys/unit-test/entities"
	"golang-journeys/unit-test/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repositories.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestUserService_GetNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", "1").Return(nil)

	user, err := userService.Get("1")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_GetSuccess(t *testing.T) {
	user := entities.User{
		Id:   "2",
		Name: "Wahyu",
	}

	userRepository.Mock.On("FindById", "2").Return(user)

	result, err := userService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Id, user.Id)
	assert.Equal(t, result.Name, user.Name)
}

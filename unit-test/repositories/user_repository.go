package repositories

import "golang-journeys/unit-test/entities"

type UserRepository interface {
	FindById(id string) *entities.User
}

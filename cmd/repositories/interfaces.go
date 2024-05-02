package repositories

import "fitness-api/cmd/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
}

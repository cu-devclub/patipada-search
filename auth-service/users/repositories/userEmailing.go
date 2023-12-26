package repositories

import (
	"auth-service/users/entities"
)

type UserEmailing interface {
	SendEmail(in *entities.Email) error 
}
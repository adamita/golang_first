package services

import (
	"github.com/adamita/first/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}

package services

import (
	"github.com/adamita/first/domain"
	"github.com/adamita/first/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}

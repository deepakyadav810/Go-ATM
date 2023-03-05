package services

import "atm-machine.com/atm-apis/models"

type UserService interface {
	CreateUser(*models.User) (int, error)
	DepositWithdraw(user []string) error
	ChangePin(user []string) error
	GetTransacion(user string) (*models.User, error)
}

package svc

import repository "github.com/trademate-be/internal/db/sqlc"

type UsersService interface {
	CreateUser(user repository.CreateUserParams) (repository.User, error)
	GetUser(userID int64) (repository.User, error)
	ListUsers(listUsersParam repository.ListUsersParams) ([]repository.User, error)
	UpdateUser(updateUserParam repository.UpdateUserParams) (repository.User, error)
}

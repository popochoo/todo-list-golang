package users_postgres_repository

import "github.com/popochoo/todo-list-golang/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func usersDomainsFromModels(users []UserModel) []domain.User {
	usersDomains := make([]domain.User, len(users))

	for i, user := range users {
		usersDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}

	return usersDomains
}

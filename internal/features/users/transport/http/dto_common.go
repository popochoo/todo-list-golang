package users_transport_http

import "github.com/popochoo/todo-list-golang/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id" example:"10"`
	Version     int     `json:"version" example:"3"`
	FullName    string  `json:"full_name" example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+7 999 222 33 44"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomain(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}

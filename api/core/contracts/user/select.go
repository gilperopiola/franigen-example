package user

import (
	"franigen-example/api/core/entities"
)

type SelectUsersRequest struct {
	ID   []uint `form:"id"`
	Name string `form:"name"`

	Limit  uint
	Offset uint
}

type SelectUsersResponse struct {
	Users []*entities.User `json:"users"`
}

func NewSelectUsersResponse(users []*entities.User) *SelectUsersResponse {
	return &SelectUsersResponse{Users: users}
}

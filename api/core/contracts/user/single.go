package user

import (
	"franigen-example/api/core/entities"
)

type SingleUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewSingleUserResponse(user *entities.User) *SingleUserResponse {
	return &SingleUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}

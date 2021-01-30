package user

import (
	"franigen-example/api/core/entities"
)

type DeleteUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewDeleteResponse(user *entities.User) *DeleteUserResponse {
	return &DeleteUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}

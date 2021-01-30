package user

import (
	"franigen-example/api/core/entities"
)

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r CreateUserRequest) NewUser() *entities.User {
	return &entities.User{
		Name: r.Name,
	}
}

func (r CreateUserRequest) Check() error {
	return nil
}

type CreateUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewCreateResponse(user *entities.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}

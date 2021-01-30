package user

import (
	"franigen-example/api/core/entities"
	"franigen-example/api/core/usecases/errors"

	"github.com/jinzhu/gorm"
)

type UpdateUserRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

func (r UpdateUserRequest) NewUser() *entities.User {
	return &entities.User{
		Model: gorm.Model{ID: r.ID},
		Name:  r.Name,
	}
}

func (r UpdateUserRequest) Check() error {
	if r.ID == 0 {
		return errors.NewMissingFieldError("id")
	}
	return nil
}

type UpdateUserResponse struct {
	User entities.User `json:"user"`
}

func NewUpdateResponse(user entities.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		User: user,
	}
}

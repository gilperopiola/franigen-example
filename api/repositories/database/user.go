package database

import (
	"context"
	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/entities"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r UserRepository) Single(ctx context.Context, userID uint) (*entities.User, error) {
	var user entities.User
	err := r.DB.First(&user, userID).Error
	return &user, err
}

func (r UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r UserRepository) Delete(ctx context.Context, userID uint) error {
	return r.DB.Delete(&entities.User{}, userID).Error
}

func (r UserRepository) Select(ctx context.Context, query *contracts.SelectUsersRequest) ([]*entities.User, error) {
	users := []*entities.User{}
	db := r.DB

	if len(query.ID) > 0 {
		db = db.Where("id IN (?)", query.ID)
	}

	if len(query.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}

	db = db.Limit(query.Limit).Offset(query.Offset)

	err := db.Find(&users).Error
	return users, err
}

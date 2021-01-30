package providers

import (
	"context"
	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/entities"
)

type User interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Single(ctx context.Context, userID uint) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Delete(ctx context.Context, userID uint) error
	Select(ctx context.Context, query *contracts.SelectUsersRequest) ([]*entities.User, error)
}

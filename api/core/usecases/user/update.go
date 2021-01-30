package user

import (
	"context"
	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/providers"
	"franigen-example/api/core/usecases/errors"

	"go.uber.org/zap"
)

type Update interface {
	Execute(ctx context.Context, req contracts.UpdateUserRequest) (*contracts.UpdateUserResponse, error)
}

type UpdateImpl struct {
	User providers.User

	Logger *zap.Logger
}

func (uc UpdateImpl) Execute(ctx context.Context, req contracts.UpdateUserRequest) (*contracts.UpdateUserResponse, error) {
	user, err := uc.User.Update(ctx, req.NewUser())
	if err != nil {
		uc.Logger.Error(errors.ErrUpdatingUser.Error(), zap.Error(err))
		return nil, errors.ErrUpdatingUser
	}

	return contracts.NewUpdateResponse(*user), nil
}

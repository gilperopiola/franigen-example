package user

import (
	"context"
	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/providers"
	"franigen-example/api/core/usecases/errors"

	"go.uber.org/zap"
)

type Create interface {
	Execute(ctx context.Context, req contracts.CreateUserRequest) (*contracts.CreateUserResponse, error)
}

type CreateImpl struct {
	User providers.User

	Logger *zap.Logger
}

func (uc CreateImpl) Execute(ctx context.Context, req contracts.CreateUserRequest) (*contracts.CreateUserResponse, error) {
	user, err := uc.User.Create(ctx, req.NewUser())
	if err != nil {
		uc.Logger.Error(errors.ErrCreatingUser.Error(), zap.Error(err))
		return nil, errors.ErrCreatingUser
	}

	return contracts.NewCreateResponse(user), nil
}

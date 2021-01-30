package user

import (
	"context"
	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/providers"
	"franigen-example/api/core/usecases/errors"

	"go.uber.org/zap"
)

type Single interface {
	Execute(ctx context.Context, userID uint) (*contracts.SingleUserResponse, error)
}

type SingleImpl struct {
	User providers.User

	Logger *zap.Logger
}

func (uc SingleImpl) Execute(ctx context.Context, userID uint) (*contracts.SingleUserResponse, error) {
	user, err := uc.User.Single(ctx, userID)
	if err != nil {
		uc.Logger.Error(errors.ErrGettingUser.Error(), zap.Error(err))
		return nil, errors.ErrGettingUser
	}

	return contracts.NewSingleUserResponse(user), nil
}

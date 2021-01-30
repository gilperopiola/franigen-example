package user

import (
	"context"
	"franigen-example/api/core/providers"
	"franigen-example/api/core/usecases/errors"

	"go.uber.org/zap"
)

type Delete interface {
	Execute(ctx context.Context, userID uint) error
}

type DeleteImpl struct {
	User providers.User

	Logger *zap.Logger
}

func (uc DeleteImpl) Execute(ctx context.Context, userID uint) error {
	err := uc.User.Delete(ctx, userID)
	if err != nil {
		uc.Logger.Error(errors.ErrDeletingUser.Error(), zap.Error(err))
		return errors.ErrDeletingUser
	}

	return nil
}

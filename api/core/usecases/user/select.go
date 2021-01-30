package user

import (
	"context"
	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/providers"
	"franigen-example/api/core/usecases/errors"

	"go.uber.org/zap"
)

type Select interface {
	Execute(ctx context.Context, req contracts.SelectUsersRequest) (*contracts.SelectUsersResponse, error)
}

type SelectImpl struct {
	User providers.User

	Logger *zap.Logger
}

func (uc SelectImpl) Execute(ctx context.Context, req contracts.SelectUsersRequest) (*contracts.SelectUsersResponse, error) {
	users, err := uc.User.Select(ctx, &req)
	if err != nil {
		uc.Logger.Error(errors.ErrGettingUsers.Error(), zap.Error(err))
		return nil, errors.ErrGettingUsers
	}

	return contracts.NewSelectUsersResponse(users), nil
}

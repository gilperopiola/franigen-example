package dependencies

import (
	"log"
	"math/rand"
	"time"

	"franigen-example/api/app/entrypoints"
	userEntry "franigen-example/api/app/entrypoints/rest/user"
	"franigen-example/api/core/usecases/user"
	"franigen-example/api/infrastructure/database"
	repositories "franigen-example/api/repositories/database"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

type RestHandlerContainer struct {
	CreateUser  entrypoints.Handler
	UpdateUser  entrypoints.Handler
	SingleUser  entrypoints.Handler
	DeleteUser  entrypoints.Handler
	SelectUsers entrypoints.Handler
}

func SetUp() *RestHandlerContainer {

	rand.Seed(time.Now().Unix())
	db := database.ConnectToDatabase()
	logger := initLogger()
	defer func() {
		_ = logger.Sync()
	}()

	// repositories
	userRepo := repositories.UserRepository{DB: db}

	// use cases
	createUserUC := user.CreateImpl{User: userRepo, Logger: logger}
	singleUserUC := user.SingleImpl{User: userRepo, Logger: logger}
	updateUserUC := user.UpdateImpl{User: userRepo, Logger: logger}
	deleteUserUC := user.DeleteImpl{User: userRepo, Logger: logger}
	selectUserUC := user.SelectImpl{User: userRepo, Logger: logger}

	// rest container
	restContainer := RestHandlerContainer{}
	restContainer.CreateUser = &userEntry.Create{UseCase: createUserUC}
	restContainer.SingleUser = &userEntry.Single{UseCase: singleUserUC}
	restContainer.UpdateUser = &userEntry.Update{UseCase: updateUserUC}
	restContainer.DeleteUser = &userEntry.Delete{UseCase: deleteUserUC}
	restContainer.SelectUsers = &userEntry.Select{UseCase: selectUserUC}

	return &restContainer
}

func initLogger() *zap.Logger {
	logger, err := zapdriver.NewProductionWithCore(zapdriver.WrapCore(
		zapdriver.ReportAllErrors(true),
	))
	if err != nil {
		log.Println("failed to init logger:", err)
	}
	return logger
}

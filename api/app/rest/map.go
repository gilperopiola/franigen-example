package rest

import (
	"franigen-example/api/app/entrypoints/rest/middleware/paraval"
	"franigen-example/api/app/entrypoints/rest/middleware/tokens"
	"franigen-example/api/infrastructure/dependencies"
	"franigen-example/config"

	"franigen-example/api/app/middleware/pages"
	"time"

	"github.com/gin-gonic/gin"
)

func Map(router *gin.Engine, handlers *dependencies.RestHandlerContainer) {

	router.Use(paraval.GinValidate())

	validate := tokens.ValidateToken(getToker(), config.Get().Env)
	paginate := pages.Paginate(pages.Values{Limit: 20})

	basePath := router.Group("/franigen")

	private := basePath.Group("", validate)
	priv := private

	private.POST("/v1/user", handlers.CreateUser.Handle)
	private.PUT("/v1/user/:userID", handlers.UpdateUser.Handle)
	private.GET("/v1/user/:userID", handlers.SingleUser.Handle)
	priv.DELETE("/v1/user/:userID", handlers.DeleteUser.Handle)
	private.GET("/v1/users", paginate, handlers.SelectUsers.Handle)
}

func getToker() tokens.Toker {
	return tokens.New(config.Get().Env, config.Get().JWTSecret, time.Hour*24*30)
}

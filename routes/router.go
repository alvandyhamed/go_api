package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go_api_learn/controller/swagger_handlers"
	"go_api_learn/middleware"
)

func SetupRouter(authSwagger *swagger_handlers.AuthSwaggerHandler,
	userSwagger *swagger_handlers.UserSwaggerHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth/register", authSwagger.Register)
	r.POST("/auth/login", authSwagger.Login)
	secured := r.Group("/users", middleware.JWTAuthMiddleware())
	{
		secured.GET("", userSwagger.GetAll)
		secured.GET("/search", userSwagger.Search)
	}

	return r
}

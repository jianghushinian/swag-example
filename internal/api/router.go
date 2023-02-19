package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/jianghushinian/swag-example/docs" // 必须导入才能注册到 gin-swagger
	"github.com/jianghushinian/swag-example/internal/api/controller/v1"
	"github.com/jianghushinian/swag-example/internal/api/controller/v2"
	"github.com/jianghushinian/swag-example/internal/api/middleware"
)

func InitController(r *gin.Engine) {
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.RedocUIMiddleware()...)
	c1 := v1.NewController()
	c2 := v2.NewController()

	r.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
	apiv1 := r.Group("/api/v1")
	{
		accounts := apiv1.Group("/accounts")
		{
			accounts.GET(":id", c1.ShowAccount)
			accounts.GET("", c1.ListAccounts)
			accounts.POST("", c1.AddAccount)
			accounts.DELETE(":id", c1.DeleteAccount)
			accounts.PATCH(":id", c1.UpdateAccount)
			accounts.POST(":id/images", c1.UploadAccountImage)
		}
		admin := apiv1.Group("/admin")
		{
			admin.Use(middleware.ApiKeyAuthMiddleware())
			admin.POST("/auth", c1.Auth) // ApiKeyAuth
		}
		examples := apiv1.Group("/examples")
		{
			examples.GET("ping", c1.PingExample)
			examples.GET("calc", c1.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", c1.PathParamsExample)
			examples.GET("header", c1.HeaderExample)
			examples.POST("security/basicauth", middleware.BasicAuthMiddleware(),
				gin.BasicAuth(gin.Accounts{
					"username": "password",
				}), c1.SecurityBasicAuthExample)
			examples.GET("attribute", c1.AttributeExample)
		}
	}

	r.GET("/swagger/v2/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2")))
	apiv2 := r.Group("/api/v2")
	{
		examples := apiv2.Group("/examples")
		{
			examples.GET("ping", c2.PingExample)
			examples.GET("calc", c2.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", c2.PathParamsExample)
			examples.GET("header", c2.HeaderExample)
			examples.GET("attribute", c2.AttributeExample)
		}
	}
}

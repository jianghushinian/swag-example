package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/mvrilo/go-redoc"
	ginRedoc "github.com/mvrilo/go-redoc/gin"

	"github.com/jianghushinian/swag-example/pkg/httputil"
)

// BasicAuthMiddleware BasicAuth 认证中间件
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("BasicAuthMiddleware Authorization: %s\n", c.GetHeader("Authorization"))
		c.Next()
	}
}

// ApiKeyAuthMiddleware ApiKeyAuth 认证中间件
func ApiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("'Authorization' is required Header"))
			c.Abort()
		}
		fmt.Printf("ApiKeyAuthMiddleware Authorization: %s\n", c.GetHeader("Authorization"))
		c.Next()
	}
}

// CorsMiddleware 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, PUT, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, jweToken, X-Requested-With",
		ExposedHeaders:  "",
		MaxAge:          60 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	})
}

// RedocUIMiddleware Redoc 风格 UI 中间件
func RedocUIMiddleware() []gin.HandlerFunc {
	_, filename, _, _ := runtime.Caller(2)
	basePath := path.Dir(path.Dir(filename))
	doc1 := redoc.Redoc{
		Title:       "ReDoc Example API v1",
		Description: "ReDoc Example API Description",
		SpecFile:    path.Join(basePath, "docs/v1_swagger.yaml"),
		SpecPath:    "/v1_swagger.yaml",
		DocsPath:    "/redoc/v1",
	}
	doc2 := redoc.Redoc{
		Title:       "ReDoc Example API v2",
		Description: "ReDoc Example API Description",
		SpecFile:    path.Join(basePath, "docs/v2_swagger.yaml"),
		SpecPath:    "/v2_swagger.yaml",
		DocsPath:    "/redoc/v2",
	}
	return []gin.HandlerFunc{ginRedoc.New(doc1), ginRedoc.New(doc2)}
}

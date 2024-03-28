package api

import (
	_ "api_exam/api/docs"
	"api_exam/api/handlers/token"
	v1 "api_exam/api/handlers/v1"
	"api_exam/config"
	"api_exam/pkg/logger"
	"api_exam/services"

	"github.com/casbin/casbin/v2"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	CasbinEnforcer *casbin.Enforcer
}

// New @Title Microservices architecture example
// @Security Auth
// @Version 1.0
// @Description This is an example of Social Network
// @Host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SigningKey,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		JWTHandler:     jwtHandler,
		Enforcer:       option.CasbinEnforcer,
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middleware.NewAuthorizer(option.CasbinEnforcer, jwtHandler, option.Conf))

	apiV1 := router.Group("/v1")

	// Users...
	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUserById)
	apiV1.GET("/users", handlerV1.GetAllUser)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id", handlerV1.DeleteUser)

	// Products...
	apiV1.POST("/products", handlerV1.CreateProduct)
	apiV1.GET("/products/by/:id", handlerV1.GetProductById)
	apiV1.GET("/products", handlerV1.GetAllProduct)
	apiV1.PUT("/products/:id", handlerV1.UpdateProduct)
	apiV1.DELETE("/products/:id", handlerV1.DeleteProduct)

	// Login...
	apiV1.POST("/register", handlerV1.Register)
	apiV1.GET("/login", handlerV1.LogIn)
	apiV1.GET("/verification", handlerV1.Verification)
	apiV1.GET("/token", handlerV1.RefreshUserToken)

	// Casbin...
	apiV1.GET("/rbac/list-role-policies", handlerV1.ListPolicies)
	apiV1.GET("/rbac/roles", handlerV1.ListRoles)
	apiV1.POST("/rbac/add-user-role", handlerV1.CreateRole)

	url := ginSwagger.URL("swagger/doc.json")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

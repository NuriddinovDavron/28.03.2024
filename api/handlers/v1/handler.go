package v1

import (
	"api_exam/api/handlers/token"
	"api_exam/config"
	"api_exam/pkg/logger"
	"api_exam/services"

	"github.com/casbin/casbin/v2"
)

type HandlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwthandler     token.JWTHandler
	enforcer       *casbin.Enforcer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     token.JWTHandler
	Enforcer       *casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwthandler:     c.JWTHandler,
		enforcer:       c.Enforcer,
	}
}

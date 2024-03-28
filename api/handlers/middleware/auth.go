package middleware

import (
	"log"
	"net/http"

	jWT "api_exam/api/handlers/token"
	"api_exam/config"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
)

type JWTRoleAuth struct {
	enforcer   *casbin.Enforcer
	cfg        config.Config
	jwtHandler jWT.JWTHandler
}

func NewAuthorizer(e *casbin.Enforcer, jwtHandler jWT.JWTHandler, cfg config.Config) gin.HandlerFunc {
	a := &JWTRoleAuth{
		enforcer:   e,
		cfg:        cfg,
		jwtHandler: jwtHandler,
	}

	return func(c *gin.Context) {
		allow, err := a.CheckPermission(c.Request)
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			if v.Errors == jwt.ValidationErrorExpired {
				a.RequireRefresh(c)
			} else {
				a.RequirePermission(c)
			}
		} else if !allow {
			a.RequirePermission(c)
		}
	}
}

func (a *JWTRoleAuth) CheckPermission(r *http.Request) (bool, error) {
	user, err := a.GetRole(r)
	if err != nil {
		return false, err
	}

	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		log.Println("failed to check permission: ", err)
		return false, err
	}

	return allowed, nil
}

func (a *JWTRoleAuth) GetRole(r *http.Request) (string, error) {
	var (
		role   string
		claims jwt.MapClaims
		err    error
	)

	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		return "unauthorized", nil
	}

	a.jwtHandler.Token = jwtToken
	claims, err = a.jwtHandler.ExtractClaims()
	if err != nil {
		return "", err
	}

	if cast.ToString(claims["role"]) == "admin" {
		role = "admin"
	} else if cast.ToString(claims["role"]) == "user" {
		role = "user"
	} else if cast.ToString(claims["role"]) == "unauthorized" {
		role = "unauthorized"
	} else {
		role = "unknown"
	}

	return role, nil
}

func (a *JWTRoleAuth) RequireRefresh(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "required refresh",
	})
	c.AbortWithStatus(401)
}

func (a *JWTRoleAuth) RequirePermission(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"Error": "You have no access this page",
	})
	c.AbortWithStatus(403)
}

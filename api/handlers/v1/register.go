package v1

import (
	"api_exam/api/handlers/models"
	"api_exam/api/handlers/token"
	"api_exam/config"
	pbu "api_exam/genproto/user_exam"
	"api_exam/pkg/etc"
	l "api_exam/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Register
// @Summary Register
// @Description Register - Api for registering users
// @Tags register
// @Accept json
// @Produce json
// @Param register body models.UserDetail true "UserDetail"
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/register [post]
func (h *HandlerV1) Register(c *gin.Context) {
	var (
		body        models.UserDetail
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind json",
		})
		h.log.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect email or password for validation")
		h.log.Error("Error validation", l.Error(err))
		return
	}

	response, err := h.serviceManager.UserService().CheckField(
		ctx, &pbu.CheckUser{
			Field: "email",
			Value: body.Email,
		})
	if err != nil || response.Exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email already use",
		})
		h.log.Error(err.Error())
		return
	}

	code := etc.GenerateCode(6)
	fmt.Println(code)
	err = etc.SetRedis(code, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "What is wrong. Please try again",
		})
	}
	etc.SendCode(body.Email, code)

	responseMessage := models.ResponseMessage{
		Content: "We send verification password you email",
	}

	c.JSON(http.StatusOK, responseMessage)
}

// LogIn
// @Summary LogIn User
// @Description LogIn - Api for login users
// @Tags register
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/login [get]
func (h *HandlerV1) LogIn(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	responseUser, err := h.serviceManager.UserService().GetUserByEmail(ctx, &pbu.EmailRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect email or password. Please try again",
		})
		h.log.Error(err.Error())
		return
	}
	if !etc.CheckPasswordHash(password, responseUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password. Please try again",
		})
		h.log.Error("failed to check password", l.Error(err))
		return
	}
	cfg := config.Load()
	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       responseUser.Email,
		Iss:       time.Now().String(),
		Exp:       cast.ToString(h.cfg.AccessTokenTimeout),
		Role:      "user",
		SigninKey: cfg.SigningKey,
		Timeout:   h.cfg.AccessTokenTimeout,
	}
	access, _, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error generating token")
		return
	}
	// parse another type with fields
	response := models.ParseStruct(responseUser, access)

	c.JSON(http.StatusOK, response)
}

// Verification
// @Summary Verification User
// @Description LogIn - Api for verification users
// @Tags register
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param code query string true "Code"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/verification [get]
func (h *HandlerV1) Verification(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	code := c.Query("code")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	userDetail, err := etc.GetRedis(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your code is expired",
		})
	}
	if email != userDetail.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect email. Try again",
		})
		return
	}

	cfg := config.Load()
	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       userDetail.Email,
		Iss:       time.Now().String(),
		Exp:       cast.ToString(h.cfg.AccessTokenTimeout),
		Role:      "user",
		SigninKey: cfg.SigningKey,
		Timeout:   h.cfg.AccessTokenTimeout,
	}
	access, _, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error generating token")
		return
	}

	createdUser, err := h.serviceManager.UserService().CreateUser(ctx, &pbu.CreateUserRequest{
		FirstName: userDetail.FirstName,
		LastName:  userDetail.LastName,
		Email:     userDetail.Email,
		Password:  userDetail.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating user",
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	response := &models.UserResponse{
		Id:           uuid.New().String(),
		FirstName:    createdUser.FirstName,
		LastName:     createdUser.LastName,
		Email:        createdUser.Email,
		Password:     createdUser.Password,
		AccessToken:  access,
		RefreshToken: createdUser.RefreshToken,
		CreatedAt:    createdUser.CreatedAt,
		UpdatedAt:    createdUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// RefreshUserToken
// @Summary Refresh User Token
// @Description This Api refresh user token
// @Tags register
// @Accept json
// @Produce json
// @Param refreshToken query string true "Refresh Token"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/token [get]
func (h *HandlerV1) RefreshUserToken(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	refreshToken := c.Query("refreshToken")

	user, err := h.serviceManager.UserService().GetUserByRefreshToken(context.Background(), &pbu.UserToken{
		RefreshToken: refreshToken,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token",
		})
		h.log.Error(err.Error())
		return
	}

	cfg := config.Load()
	// generate JWT token
	h.jwthandler = token.JWTHandler{
		Sub:       user.Email,
		Iss:       time.Now().String(),
		Exp:       cast.ToString(h.cfg.AccessTokenTimeout),
		Role:      "user",
		SigninKey: cfg.SigningKey,
		Timeout:   h.cfg.AccessTokenTimeout,
	}
	access, _, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error generating token")
		return
	}
	// update user
	updateUser, err := h.serviceManager.UserService().UpdateUser(context.Background(), &pbu.User{
		RefreshToken: user.RefreshToken,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "What is wrong. Please try again",
		})
		h.log.Error(err.Error())
		return
	}
	// parse another type with fields
	response := models.ParseStruct(updateUser, access)

	c.JSON(http.StatusOK, response)
}

package v1

import (
	"api_exam/pkg/etc"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"api_exam/api/handlers/models"
	pbu "api_exam/genproto/user_exam"
	l "api_exam/pkg/logger"
	"api_exam/pkg/utils"
)

// CreateUser ...
// @Summary CreateUser
// @Description Api for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.CreateUser true "CreateUser"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/ [post]
func (h *HandlerV1) CreateUser(c *gin.Context) {
	var (
		body        models.CreateUser
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	hashPassword, err := etc.HashPassword(body.Password)
	if err != nil {
		return
	}

	_, refresh, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		return
	}


	//mock__________
	response, err := h.serviceManager.MockUserService().CreateUser(ctx, pbu.CreateUserRequest{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Email:        body.Email,
		Password:     hashPassword,
		RefreshToken: refresh,
	})
	//mock_________end

	// response, err := h.serviceManager.UserService().CreateUser(ctx, &pbu.CreateUserRequest{
	// 	FirstName:    body.FirstName,
	// 	LastName:     body.LastName,
	// 	Email:        body.Email,
	// 	Password:     hashPassword,
	// 	RefreshToken: refresh,
	// })
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUserById GetUser gets user by id
// @Summary GetUser
// @Description Api for getting user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/{id} [get]
func (h *HandlerV1) GetUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_______
	response, err := h.serviceManager.MockUserService().GetUserById(
		ctx, &pbu.GetUserByIdRequest{
			UserId: id,
		})
	//mock________end

	// response, err := h.serviceManager.UserService().GetUserById(
	// 	ctx, &pbu.GetUserByIdRequest{
	// 		UserId: id,
	// 	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllUser ListUsers returns list of users
// @Summary ListUser
// @Description Api returns list of users
// @Tags user
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/ [get]
func (h *HandlerV1) GetAllUser(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_________
	response, err := h.serviceManager.MockUserService().GetAllUser(
		ctx, &pbu.GetAllUserRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	//mock________end

	// response, err := h.serviceManager.UserService().GetAllUser(
	// 	ctx, &pbu.GetAllUserRequest{
	// 		Limit: params.Limit,
	// 		Page:  params.Page,
	// 	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by id
// @Summary UpdateUser
// @Description Api returns updates user
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.User true "updatedUserModel"
// @Success 200 {Object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/{id} [put]
func (h *HandlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pbu.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock__________
	response, err := h.serviceManager.MockUserService().UpdateUser(ctx, &body)
	//mock__________end

	// response, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes user by id
// @Summary DeleteUser
// @Description Api deletes user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {Object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/{id} [delete]
func (h *HandlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_________
	err := h.serviceManager.MockUserService().DeleteUser(
			ctx, &pbu.GetUserByIdRequest{
				UserId: guid,
			})
	//mock_________end

	// response, err := h.serviceManager.UserService().DeleteUser(
	// 	ctx, &pbu.GetUserByIdRequest{
	// 		UserId: guid,
	// 	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, err)
}

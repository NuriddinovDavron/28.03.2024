package v1

import (
	"api_exam/api/handlers/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListPolicies @Security      ApiKeyAuth
// @Summary       Get list of policies
// @Description   This API get list of policies
// @Tags          casbin
// @Accept        json
// @Produce       json
// @Param         role query string true "Role"
// @Success        200 {object} models.ListPolePolicyResponse
// @Failure       404 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/rbac/list-role-policies [GET]
func (h *HandlerV1) ListPolicies(ctx *gin.Context) {
	role := ctx.Query("role")
	var resp models.ListPolePolicyResponse

	for _, p := range h.enforcer.GetFilteredPolicy(0, role) {
		resp.Policies = append(resp.Policies, &models.Policy{
			Role:     p[0],
			Endpoint: p[1],
			Method:   p[2],
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

// ListRoles @Security      ApiKeyAuth
// @Summary       Get list of roles
// @Description   This API get list of roles
// @Tags          casbin
// @Accept        json
// @Produce       json
// @Param         limit query int false "limit"
// @Param         offset query int false "offset"
// @Success        200 {object} []string
// @Failure       404 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/rbac/roles [GET]
func (h *HandlerV1) ListRoles(ctx *gin.Context) {
	resp := h.enforcer.GetAllRoles()
	ctx.JSON(http.StatusOK, resp)
}

// CreateRole
// @Security      ApiKeyAuth
// @Router        /v1/rbac/add-user-role [POST]
// @Summary       Create new user-role
// @Description   Create new user-role
// @Tags          casbin
// @Accept        json
// @Produce       json
// @Param         Create body  models.CreateUserRoleRequest true  "body"
// @Success       200 {object} models.CreateUserRoleRequest
// @Failure 	  404 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router        /v1/rbac/add-user-role [POST]
func (h *HandlerV1) CreateRole(ctx *gin.Context) {
	var reqBody models.CreateUserRoleRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody); err != nil {
		h.log.Error("rbacHandler/CreateUserRole", zap.Error(err))
		return
	}

	// h.enforcer.AddPolicy()
	if _, err := h.enforcer.AddRoleForUser(reqBody.UserId, reqBody.Path); err != nil {
		h.log.Error("error on grantAccess", zap.Error(err))
		return
	}
	h.enforcer.SavePolicy()
	ctx.JSON(http.StatusOK, reqBody)
}

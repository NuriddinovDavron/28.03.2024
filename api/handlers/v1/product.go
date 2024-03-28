package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"api_exam/api/handlers/models"
	pbp "api_exam/genproto/product_exam"
	l "api_exam/pkg/logger"
	"api_exam/pkg/utils"
)

// CreateProduct ...
// @Summary CreateProduct
// @Description Api for creating a new product
// @Tags product
// @Accept json
// @Produce json
// @Param Product body models.Product true "createProductModel"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/products/ [post]
func (h *HandlerV1) CreateProduct(c *gin.Context) {
	var (
		body        models.Product
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

	//mock___________
	response, err := h.serviceManager.MockProductService().CreateProduct(ctx, &pbp.CreateProductRequest{
		Name:        body.Name,
		Description: body.Description,
		OwnerId:     body.OwnerId,
		Price:       int64(body.Price),
	})
	//mock________end

	// response, err := h.serviceManager.ProductService().CreateProduct(ctx, &pbp.CreateProductRequest{
	// 	Name:        body.Name,
	// 	Description: body.Description,
	// 	OwnerId:     body.OwnerId,
	// 	Price:       int64(body.Price),
	// })
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create product", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetProductById GetProduct gets product by id
// @Summary GetProduct
// @Description Api for getting product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/products/by/{id} [get]
func (h *HandlerV1) GetProductById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// mock___________
	response, err := h.serviceManager.MockProductService().GetProductById(
		ctx, &pbp.GetProductByIdRequest{
			ProductId: id,
		})
	// mock_________end

	// response, err := h.serviceManager.ProductService().GetProductById(
	// 	ctx, &pbp.GetProductByIdRequest{
	// 		ProductId: id,
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

// GetAllProduct ListProducts returns list of products
// @Summary ListProducts
// @Description Api returns list of products
// @Tags product
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/products/ [get]
func (h *HandlerV1) GetAllProduct(c *gin.Context) {
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

	//mock_______
	response, err := h.serviceManager.MockProductService().GetAllProduct(
		ctx, &pbp.GetAllProductRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	//mock________end

	// response, err := h.serviceManager.ProductService().GetAllProduct(
	// 	ctx, &pbp.GetAllProductRequest{
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

// UpdateProduct updates product by id
// @Summary UpdateProduct
// @Description Api returns updated product
// @Tags product
// @Accept json
// @Produce json
// @Param Product body models.Product true "createProductModel"
// @Success 200 {Object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/products/{id} [put]
func (h *HandlerV1) UpdateProduct(c *gin.Context) {
	var (
		body        pbp.UpdateProductRequest
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

	//mock_________
	response, err := h.serviceManager.MockProductService().UpdateProduct(ctx, &body)
	//mock_______end

	// response, err := h.serviceManager.ProductService().UpdateProduct(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProduct deletes product by id
// @Summary DeleteProduct
// @Description Api deletes product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {Object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/products/{id} [delete]
func (h *HandlerV1) DeleteProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_______
	 err := h.serviceManager.MockProductService().DeleteProduct(
		ctx, &pbp.GetProductByIdRequest{
			ProductId: id,
		})
	//mock________end

	// response, err := h.serviceManager.ProductService().DeleteProduct(
	// 	ctx, &pbp.GetProductByIdRequest{
	// 		ProductId: id,
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

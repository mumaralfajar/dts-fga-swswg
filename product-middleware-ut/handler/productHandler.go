package handler

import (
	"dts-fga-swswg/product-middleware-ut/dto"
	"dts-fga-swswg/product-middleware-ut/entity"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
	"dts-fga-swswg/product-middleware-ut/pkg/helpers"
	"dts-fga-swswg/product-middleware-ut/service"
	"net/http"

	_ "dts-fga-swswg/product-middleware-ut/entity"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

// CreateNewProduct godoc
// @Tags products
// @Description Create New Product Data
// @ID create-new-product
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewProductRequest true "request body json"
// @Success 201 {object} dto.NewProductRequest
// @Router /products [post]
func (m productHandler) CreateProduct(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newProduct, err := m.productService.CreateProduct(user.Id, productRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (m productHandler) UpdateProductById(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	productId, err := helpers.GetParamId(c, "productId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.productService.UpdateProductById(productId, productRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

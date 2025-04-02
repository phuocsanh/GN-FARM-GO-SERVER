package product

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/service"
	"go_ecommerce/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController manages product-related endpoints
var Product = new(cProduct)

type cProduct struct{}

// CreateProduct creates a new product
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags product management
// @Accept json
// @Produce json
// @Param payload body model.ProductInput true "Product details"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/create [post]
func (c *cProduct) CreateProduct(ctx *gin.Context) {
	var input model.ProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	product, err := service.ProductManagement().CreateProduct(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, product)
}

// UpdateProduct updates an existing product
// @Summary Update an existing product
// @Description Update an existing product with the provided details
// @Tags product management
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param payload body model.ProductInput true "Product details"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/update/{id} [put]
func (c *cProduct) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if productID == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Product ID is required")
		return
	}

	var input model.ProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	err := service.ProductManagement().UpdateProduct(ctx, productID, &input)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}

// PublishProduct publishes a product
// @Summary Publish a product
// @Description Change a product's status to published
// @Tags product management
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/publish/{id} [put]
func (c *cProduct) PublishProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if productID == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Product ID is required")
		return
	}

	userID, _ := ctx.Get("user_id")
	shopID := userID.(string)

	err := service.ProductManagement().PublishProduct(ctx, productID, shopID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}

// UnPublishProduct unpublishes a product
// @Summary Unpublish a product
// @Description Change a product's status to unpublished
// @Tags product management
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/unpublish/{id} [put]
func (c *cProduct) UnPublishProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if productID == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Product ID is required")
		return
	}

	userID, _ := ctx.Get("user_id")
	shopID := userID.(string)

	err := service.ProductManagement().UnPublishProduct(ctx, productID, shopID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}

// GetProductByID gets a product by ID
// @Summary Get a product by ID
// @Description Get detailed information about a product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/{id} [get]
func (c *cProduct) GetProductByID(ctx *gin.Context) {
	productID := ctx.Param("id")
	if productID == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Product ID is required")
		return
	}

	product, err := service.ProductManagement().FindProduct(ctx, productID)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, product)
}

// GetAllProducts gets all products
// @Summary Get all products
// @Description Get a list of products with optional filtering
// @Tags product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param product_type query string false "Product type filter"
// @Param keyword query string false "Search keyword"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product [get]
func (c *cProduct) GetAllProducts(ctx *gin.Context) {
	params := model.ProductQueryParams{}
	
	// Parse page parameter
	page := ctx.DefaultQuery("page", "1")
	if pageNum, err := strconv.Atoi(page); err == nil {
		params.Page = pageNum
	} else {
		params.Page = 1
	}
	
	// Parse limit parameter
	limit := ctx.DefaultQuery("limit", "10")
	if limitNum, err := strconv.Atoi(limit); err == nil {
		params.Limit = limitNum
	} else {
		params.Limit = 10
	}
	
	// Get other filters
	params.ProductType = ctx.Query("product_type")
	params.Keyword = ctx.Query("keyword")

	products, err := service.ProductManagement().FindAllProducts(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, products)
}

// GetAllDraftsForShop gets all draft products for a shop
// @Summary Get all draft products for a shop
// @Description Get a list of draft products for the current shop
// @Tags product management
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/drafts [get]
func (c *cProduct) GetAllDraftsForShop(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	shopID := userID.(string)
	
	// Parse page parameter
	page := ctx.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		pageNum = 1
	}
	
	// Parse limit parameter
	limit := ctx.DefaultQuery("limit", "10")
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		limitNum = 10
	}

	products, err := service.ProductManagement().FindAllDraftsForShop(ctx, shopID, pageNum, limitNum)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, products)
}

// GetAllPublishForShop gets all published products for a shop
// @Summary Get all published products for a shop
// @Description Get a list of published products for the current shop
// @Tags product management
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/published [get]
func (c *cProduct) GetAllPublishForShop(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	shopID := userID.(string)
	
	// Parse page parameter
	page := ctx.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		pageNum = 1
	}
	
	// Parse limit parameter
	limit := ctx.DefaultQuery("limit", "10")
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		limitNum = 10
	}

	products, err := service.ProductManagement().FindAllPublishForShop(ctx, shopID, pageNum, limitNum)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, products)
}

// GetProductsByDiscount gets products ordered by discount
// @Summary Get products ordered by discount
// @Description Get a list of products ordered by discount amount
// @Tags product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param product_type query string false "Product type filter"
// @Param keyword query string false "Search keyword"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/discounts [get]
func (c *cProduct) GetProductsByDiscount(ctx *gin.Context) {
	params := model.ProductQueryParams{}
	
	// Parse page parameter
	page := ctx.DefaultQuery("page", "1")
	if pageNum, err := strconv.Atoi(page); err == nil {
		params.Page = pageNum
	} else {
		params.Page = 1
	}
	
	// Parse limit parameter
	limit := ctx.DefaultQuery("limit", "10")
	if limitNum, err := strconv.Atoi(limit); err == nil {
		params.Limit = limitNum
	} else {
		params.Limit = 10
	}
	
	// Get other filters
	params.ProductType = ctx.Query("product_type")
	params.Keyword = ctx.Query("keyword")

	products, err := service.ProductManagement().GetProductsByDiscount(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, products)
}

// GetProductsBySelled gets products ordered by number sold
// @Summary Get products ordered by number sold
// @Description Get a list of products ordered by number sold
// @Tags product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param product_type query string false "Product type filter"
// @Param keyword query string false "Search keyword"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/bestsellers [get]
func (c *cProduct) GetProductsBySelled(ctx *gin.Context) {
	params := model.ProductQueryParams{}
	
	// Parse page parameter
	page := ctx.DefaultQuery("page", "1")
	if pageNum, err := strconv.Atoi(page); err == nil {
		params.Page = pageNum
	} else {
		params.Page = 1
	}
	
	// Parse limit parameter
	limit := ctx.DefaultQuery("limit", "10")
	if limitNum, err := strconv.Atoi(limit); err == nil {
		params.Limit = limitNum
	} else {
		params.Limit = 10
	}
	
	// Get other filters
	params.ProductType = ctx.Query("product_type")
	params.Keyword = ctx.Query("keyword")

	products, err := service.ProductManagement().GetProductsBySelled(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, products)
}

// SearchProducts searches products by keyword
// @Summary Search products by keyword
// @Description Search for products using a keyword
// @Tags product
// @Accept json
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /product/search [get]
func (c *cProduct) SearchProducts(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	if keyword == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Keyword is required")
		return
	}

	products, err := service.ProductManagement().SearchProducts(ctx, keyword)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, products)
} 
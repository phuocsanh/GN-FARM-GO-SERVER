package impl

import (
	"context"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/repo"
	"go_ecommerce/internal/service"
	"go_ecommerce/internal/utils/auth"
	"strings"
	"time"

	"github.com/google/uuid"
)

type productService struct {
	productRepo repo.IProductRepository
}

// NewProductService tạo một instance mới của service product
func NewProductService() service.IProductManagement {
	return &productService{
		productRepo: repo.NewProductRepository(),
	}
}

// Đảm bảo productService implement interface IProductManagement
var _ service.IProductManagement = (*productService)(nil)

// CreateProduct tạo một sản phẩm mới dựa trên loại sản phẩm
func (s *productService) CreateProduct(ctx context.Context, input *model.ProductInput) (interface{}, error) {
	// Validate input
	if input.ProductName == "" || input.ProductPrice <= 0 {
		return nil, ErrInvalidInput
	}

	// Get user from context or auth token
	userId, err := auth.ExtractUserID(ctx)
	if err != nil {
		return nil, err
	}

	// Process HTML description
	description := input.ProductDescription
	if description != "" {
		htmlContent := ""
		lines := strings.Split(description, "\n")
		for _, line := range lines {
			if line = strings.TrimSpace(line); line != "" {
				htmlContent += "<p>" + line + "</p>"
			}
		}
		input.ProductDescription = htmlContent
	}

	// Generate product ID
	productID := uuid.New().String()

	// Create main product record
	product := &model.ProductModel{
		ID:                   productID,
		ProductName:          input.ProductName,
		ProductPrice:         input.ProductPrice,
		ProductDiscountPrice: input.ProductDiscountPrice,
		ProductThumb:         input.ProductThumb,
		ProductDescription:   input.ProductDescription,
		ProductQuantity:      input.ProductQuantity,
		ProductType:          input.ProductType,
		SubProductType:       input.SubProductType,
		ProductVideos:        input.ProductVideos,
		ProductPictures:      input.ProductPictures,
		ProductStatus:        input.ProductStatus,
		ProductShop:          userId,
		IsDraft:              true,
		IsPublished:          false,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	// Create specific product type based on product_type
	var err2 error
	switch input.ProductType {
	case "Mushroom":
		mushroom := &model.MushroomModel{
			ID:          productID,
			ProductShop: userId,
		}

		// Extract attributes from input
		if attrs, ok := input.ProductAttributes["weight"].(float64); ok {
			mushroom.Weight = attrs
		}
		if attrs, ok := input.ProductAttributes["origin"].(string); ok {
			mushroom.Origin = attrs
		}
		if attrs, ok := input.ProductAttributes["freshness"].(string); ok {
			mushroom.Freshness = attrs
		}
		if attrs, ok := input.ProductAttributes["package_type"].(string); ok {
			mushroom.PackageType = attrs
		}

		err2 = s.productRepo.CreateMushroom(ctx, mushroom)
	case "Vegetable":
		vegetable := &model.VegetableModel{
			ID:          productID,
			ProductShop: userId,
		}

		// Extract attributes from input
		if attrs, ok := input.ProductAttributes["weight"].(float64); ok {
			vegetable.Weight = attrs
		}
		if attrs, ok := input.ProductAttributes["origin"].(string); ok {
			vegetable.Origin = attrs
		}
		if attrs, ok := input.ProductAttributes["freshness"].(string); ok {
			vegetable.Freshness = attrs
		}
		if attrs, ok := input.ProductAttributes["package_type"].(string); ok {
			vegetable.PackageType = attrs
		}

		err2 = s.productRepo.CreateVegetable(ctx, vegetable)
	case "Bonsai":
		bonsai := &model.BonsaiModel{
			ID:          productID,
			ProductShop: userId,
		}

		// Extract attributes from input
		if attrs, ok := input.ProductAttributes["age"].(int); ok {
			bonsai.Age = attrs
		}
		if attrs, ok := input.ProductAttributes["height"].(int); ok {
			bonsai.Height = attrs
		}
		if attrs, ok := input.ProductAttributes["style"].(string); ok {
			bonsai.Style = attrs
		}
		if attrs, ok := input.ProductAttributes["species"].(string); ok {
			bonsai.Species = attrs
		}
		if attrs, ok := input.ProductAttributes["pot_type"].(string); ok {
			bonsai.PotType = attrs
		}

		err2 = s.productRepo.CreateBonsai(ctx, bonsai)
	default:
		return nil, ErrInvalidProductType
	}

	if err2 != nil {
		return nil, err2
	}

	// Create the main product
	err = s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	// Create inventory record
	inventory := &model.InventoryInput{
		ProductID: productID,
		ShopID:    userId,
		Stock:     input.ProductQuantity,
	}
	err = s.productRepo.InsertInventory(ctx, inventory)
	if err != nil {
		// Handle error, maybe delete the product if inventory creation fails
		return nil, err
	}

	return product, nil
}

// UpdateProduct cập nhật sản phẩm hiện có
func (s *productService) UpdateProduct(ctx context.Context, productID string, input *model.ProductInput) error {
	userId, err := auth.ExtractUserID(ctx)
	if err != nil {
		return err
	}

	// Find the product to check ownership
	product, err := s.productRepo.FindProduct(ctx, productID)
	if err != nil {
		return err
	}

	if product.ProductShop != userId {
		return ErrUnauthorized
	}

	// Prepare update data
	updateData := map[string]interface{}{}

	if input.ProductName != "" {
		updateData["product_name"] = input.ProductName
	}
	if input.ProductPrice > 0 {
		updateData["product_price"] = input.ProductPrice
	}
	if input.ProductThumb != "" {
		updateData["product_thumb"] = input.ProductThumb
	}
	if input.ProductDescription != "" {
		// Process HTML description
		description := input.ProductDescription
		htmlContent := ""
		lines := strings.Split(description, "\n")
		for _, line := range lines {
			if line = strings.TrimSpace(line); line != "" {
				htmlContent += "<p>" + line + "</p>"
			}
		}
		updateData["product_description"] = htmlContent
	}
	if input.ProductQuantity > 0 {
		updateData["product_quantity"] = input.ProductQuantity
	}
	if input.ProductStatus != "" {
		updateData["product_status"] = input.ProductStatus
	}
	if input.ProductDiscountPrice > 0 {
		updateData["product_discounted_price"] = input.ProductDiscountPrice
	}
	if len(input.ProductVideos) > 0 {
		updateData["product_videos"] = input.ProductVideos
	}
	if len(input.ProductPictures) > 0 {
		updateData["product_pictures"] = input.ProductPictures
	}
	if input.SubProductType != "" {
		updateData["sub_product_type"] = input.SubProductType
	}

	updateData["updated_at"] = time.Now()

	// Update specific product type attributes
	switch product.ProductType {
	case "Mushroom":
		mushroomAttrs := map[string]interface{}{}
		if attrs, ok := input.ProductAttributes["weight"].(float64); ok {
			mushroomAttrs["weight"] = attrs
		}
		if attrs, ok := input.ProductAttributes["origin"].(string); ok {
			mushroomAttrs["origin"] = attrs
		}
		if attrs, ok := input.ProductAttributes["freshness"].(string); ok {
			mushroomAttrs["freshness"] = attrs
		}
		if attrs, ok := input.ProductAttributes["package_type"].(string); ok {
			mushroomAttrs["package_type"] = attrs
		}

		if len(mushroomAttrs) > 0 {
			err = s.productRepo.UpdateProductByID(ctx, productID, mushroomAttrs)
			if err != nil {
				return err
			}
		}
	case "Vegetable":
		vegetableAttrs := map[string]interface{}{}
		if attrs, ok := input.ProductAttributes["weight"].(float64); ok {
			vegetableAttrs["weight"] = attrs
		}
		if attrs, ok := input.ProductAttributes["origin"].(string); ok {
			vegetableAttrs["origin"] = attrs
		}
		if attrs, ok := input.ProductAttributes["freshness"].(string); ok {
			vegetableAttrs["freshness"] = attrs
		}
		if attrs, ok := input.ProductAttributes["package_type"].(string); ok {
			vegetableAttrs["package_type"] = attrs
		}

		if len(vegetableAttrs) > 0 {
			err = s.productRepo.UpdateProductByID(ctx, productID, vegetableAttrs)
			if err != nil {
				return err
			}
		}
	case "Bonsai":
		bonsaiAttrs := map[string]interface{}{}
		if attrs, ok := input.ProductAttributes["age"].(int); ok {
			bonsaiAttrs["age"] = attrs
		}
		if attrs, ok := input.ProductAttributes["height"].(int); ok {
			bonsaiAttrs["height"] = attrs
		}
		if attrs, ok := input.ProductAttributes["style"].(string); ok {
			bonsaiAttrs["style"] = attrs
		}
		if attrs, ok := input.ProductAttributes["species"].(string); ok {
			bonsaiAttrs["species"] = attrs
		}
		if attrs, ok := input.ProductAttributes["pot_type"].(string); ok {
			bonsaiAttrs["pot_type"] = attrs
		}

		if len(bonsaiAttrs) > 0 {
			err = s.productRepo.UpdateProductByID(ctx, productID, bonsaiAttrs)
			if err != nil {
				return err
			}
		}
	}

	// Update main product
	if len(updateData) > 0 {
		return s.productRepo.UpdateProductByID(ctx, productID, updateData)
	}

	return nil
}

// PublishProduct đăng tải sản phẩm
func (s *productService) PublishProduct(ctx context.Context, productID string, shopID string) error {
	return s.productRepo.PublishProductByShop(ctx, productID, shopID)
}

// UnPublishProduct hủy đăng tải sản phẩm
func (s *productService) UnPublishProduct(ctx context.Context, productID string, shopID string) error {
	return s.productRepo.UnPublishProductByShop(ctx, productID, shopID)
}

// FindProduct tìm sản phẩm theo ID
func (s *productService) FindProduct(ctx context.Context, productID string) (*model.ProductModel, error) {
	return s.productRepo.FindProduct(ctx, productID)
}

// FindAllProducts tìm tất cả sản phẩm theo tham số
func (s *productService) FindAllProducts(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error) {
	return s.productRepo.FindAllProducts(ctx, params)
}

// FindAllDraftsForShop tìm tất cả sản phẩm nháp của một shop
func (s *productService) FindAllDraftsForShop(ctx context.Context, shopID string, page, limit int) ([]model.ProductModel, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	return s.productRepo.FindAllDraftsForShop(ctx, shopID, limit, offset)
}

// FindAllPublishForShop tìm tất cả sản phẩm đã đăng tải của một shop
func (s *productService) FindAllPublishForShop(ctx context.Context, shopID string, page, limit int) ([]model.ProductModel, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	return s.productRepo.FindAllPublishForShop(ctx, shopID, limit, offset)
}

// GetProductsByDiscount lấy sản phẩm sắp xếp theo giảm giá
func (s *productService) GetProductsByDiscount(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error) {
	return s.productRepo.FindProductsByDiscount(ctx, params)
}

// GetProductsBySelled lấy sản phẩm sắp xếp theo số lượng đã bán
func (s *productService) GetProductsBySelled(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error) {
	return s.productRepo.FindProductsBySelled(ctx, params)
}

// SearchProducts tìm kiếm sản phẩm theo từ khóa
func (s *productService) SearchProducts(ctx context.Context, keyword string) ([]model.ProductModel, error) {
	return s.productRepo.SearchProducts(ctx, keyword)
} 
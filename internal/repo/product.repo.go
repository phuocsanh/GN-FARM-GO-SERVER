package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go_ecommerce/global"
	"go_ecommerce/internal/database"
	"go_ecommerce/internal/model"
	"math"
	"strconv"

	"gorm.io/gorm"
)

type IProductRepository interface {
	// Create methods
	CreateProduct(ctx context.Context, product *model.ProductModel) error
	CreateMushroom(ctx context.Context, mushroom *model.MushroomModel) error
	CreateVegetable(ctx context.Context, vegetable *model.VegetableModel) error
	CreateBonsai(ctx context.Context, bonsai *model.BonsaiModel) error
	InsertInventory(ctx context.Context, inventory *model.InventoryInput) error
	
	// Read methods
	FindProduct(ctx context.Context, productID string) (*model.ProductModel, error)
	
	// Update methods
	UpdateProductByID(ctx context.Context, productID string, updateData map[string]interface{}) error
	PublishProductByShop(ctx context.Context, productID string, shopID string) error
	UnPublishProductByShop(ctx context.Context, productID string, shopID string) error
	
	// List methods
	FindAllDraftsForShop(ctx context.Context, shopID string, limit, offset int) ([]model.ProductModel, error)
	FindAllPublishForShop(ctx context.Context, shopID string, limit, offset int) ([]model.ProductModel, error)
	FindAllProducts(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error)
	FindProductsByDiscount(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error)
	FindProductsBySelled(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error)
	SearchProducts(ctx context.Context, keyword string) ([]model.ProductModel, error)
	
	// Count methods
	CountProducts(ctx context.Context, query map[string]interface{}) (int64, error)
}

type productRepository struct {
	db   *gorm.DB
	sqlc *database.Queries
}

func NewProductRepository() IProductRepository {
	return &productRepository{
		db:   global.Mdb,
		sqlc: database.New(global.Mdbc),
	}
}

// CreateProduct creates a new product using sqlc
func (p *productRepository) CreateProduct(ctx context.Context, product *model.ProductModel) error {
	// Convert JSON arrays to strings for database
	var videoJSON, pictureJSON json.RawMessage
	
	if len(product.ProductVideos) > 0 {
		videoBytes, err := json.Marshal(product.ProductVideos)
		if err == nil {
			videoJSON = videoBytes
		}
	}
	
	if len(product.ProductPictures) > 0 {
		pictureBytes, err := json.Marshal(product.ProductPictures)
		if err == nil {
			pictureJSON = pictureBytes
		}
	}
	
	// Set up discount price as nullable
	var discountPrice sql.NullString
	if product.ProductDiscountPrice > 0 {
		discountPrice = sql.NullString{
			String: fmt.Sprintf("%.2f", product.ProductDiscountPrice),
			Valid:  true,
		}
	}
	
	// Call sqlc generated function
	_, err := p.sqlc.CreateProduct(ctx, database.CreateProductParams{
		ID:                     product.ID,
		ProductName:            product.ProductName,
		ProductPrice:           fmt.Sprintf("%.2f", product.ProductPrice),
		ProductDiscountedPrice: discountPrice,
		ProductThumb:           sql.NullString{String: product.ProductThumb, Valid: product.ProductThumb != ""},
		ProductDescription:     sql.NullString{String: product.ProductDescription, Valid: product.ProductDescription != ""},
		ProductQuantity:        int32(product.ProductQuantity),
		ProductType:            product.ProductType,
		SubProductType:         sql.NullString{String: product.SubProductType, Valid: product.SubProductType != ""},
		ProductVideos:          videoJSON,
		ProductPictures:        pictureJSON,
		ProductStatus:          product.ProductStatus,
		ProductShop:            product.ProductShop,
		IsDraft:                product.IsDraft,
		IsPublished:            product.IsPublished,
	})
	
	return err
}

// CreateMushroom creates a new mushroom product using sqlc
func (p *productRepository) CreateMushroom(ctx context.Context, mushroom *model.MushroomModel) error {
	_, err := p.sqlc.CreateMushroom(ctx, database.CreateMushroomParams{
		ID:          mushroom.ID,
		ProductShop: mushroom.ProductShop,
		Weight:      sql.NullString{String: fmt.Sprintf("%.2f", mushroom.Weight), Valid: mushroom.Weight > 0},
		Origin:      sql.NullString{String: mushroom.Origin, Valid: mushroom.Origin != ""},
		Freshness:   sql.NullString{String: mushroom.Freshness, Valid: mushroom.Freshness != ""},
		PackageType: sql.NullString{String: mushroom.PackageType, Valid: mushroom.PackageType != ""},
	})
	
	return err
}

// CreateVegetable creates a new vegetable product using sqlc
func (p *productRepository) CreateVegetable(ctx context.Context, vegetable *model.VegetableModel) error {
	_, err := p.sqlc.CreateVegetable(ctx, database.CreateVegetableParams{
		ID:          vegetable.ID,
		ProductShop: vegetable.ProductShop,
		Weight:      sql.NullString{String: fmt.Sprintf("%.2f", vegetable.Weight), Valid: vegetable.Weight > 0},
		Origin:      sql.NullString{String: vegetable.Origin, Valid: vegetable.Origin != ""},
		Freshness:   sql.NullString{String: vegetable.Freshness, Valid: vegetable.Freshness != ""},
		PackageType: sql.NullString{String: vegetable.PackageType, Valid: vegetable.PackageType != ""},
	})
	
	return err
}

// CreateBonsai creates a new bonsai product using sqlc
func (p *productRepository) CreateBonsai(ctx context.Context, bonsai *model.BonsaiModel) error {
	_, err := p.sqlc.CreateBonsai(ctx, database.CreateBonsaiParams{
		ID:          bonsai.ID,
		ProductShop: bonsai.ProductShop,
		Age:         sql.NullInt32{Int32: int32(bonsai.Age), Valid: bonsai.Age > 0},
		Height:      sql.NullInt32{Int32: int32(bonsai.Height), Valid: bonsai.Height > 0},
		Style:       sql.NullString{String: bonsai.Style, Valid: bonsai.Style != ""},
		Species:     sql.NullString{String: bonsai.Species, Valid: bonsai.Species != ""},
		PotType:     sql.NullString{String: bonsai.PotType, Valid: bonsai.PotType != ""},
	})
	
	return err
}

// InsertInventory creates a new inventory entry using sqlc
func (p *productRepository) InsertInventory(ctx context.Context, inventory *model.InventoryInput) error {
	_, err := p.sqlc.CreateInventory(ctx, database.CreateInventoryParams{
		ProductID: inventory.ProductID,
		ShopID:    inventory.ShopID,
		Location:  sql.NullString{String: inventory.Location, Valid: inventory.Location != ""},
		Stock:     int32(inventory.Stock),
	})
	
	return err
}

// FindProduct finds a product by ID using sqlc
func (p *productRepository) FindProduct(ctx context.Context, productID string) (*model.ProductModel, error) {
	dbProduct, err := p.sqlc.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	
	// Convert string price to float64
	price, _ := strconv.ParseFloat(dbProduct.ProductPrice, 64)
	
	product := &model.ProductModel{
		ID:                   dbProduct.ID,
		ProductName:          dbProduct.ProductName,
		ProductPrice:         price,
		ProductThumb:         dbProduct.ProductThumb.String,
		ProductDescription:   dbProduct.ProductDescription.String,
		ProductQuantity:      int(dbProduct.ProductQuantity),
		ProductType:          dbProduct.ProductType,
		SubProductType:       dbProduct.SubProductType.String,
		ProductStatus:        dbProduct.ProductStatus,
		ProductSelled:        int(dbProduct.ProductSelled),
		ProductShop:          dbProduct.ProductShop,
		IsDraft:              dbProduct.IsDraft,
		IsPublished:          dbProduct.IsPublished,
	}
	
	// Handle nullables
	if dbProduct.CreatedAt.Valid {
		product.CreatedAt = dbProduct.CreatedAt.Time
	}
	
	if dbProduct.UpdatedAt.Valid {
		product.UpdatedAt = dbProduct.UpdatedAt.Time
	}
	
	// Handle discounted price
	if dbProduct.ProductDiscountedPrice.Valid {
		discountedPrice, _ := strconv.ParseFloat(dbProduct.ProductDiscountedPrice.String, 64)
		product.ProductDiscountPrice = discountedPrice
	}
	
	// Unmarshal JSON arrays
	if len(dbProduct.ProductVideos) > 0 {
		var videos []string
		if err := json.Unmarshal(dbProduct.ProductVideos, &videos); err == nil {
			product.ProductVideos = videos
		}
	}
	
	if len(dbProduct.ProductPictures) > 0 {
		var pictures []string
		if err := json.Unmarshal(dbProduct.ProductPictures, &pictures); err == nil {
			product.ProductPictures = pictures
		}
	}
	
	return product, nil
}

// UpdateProductByID updates a product by ID using gorm
func (p *productRepository) UpdateProductByID(ctx context.Context, productID string, updateData map[string]interface{}) error {
	return p.db.Model(&model.ProductModel{}).Where("id = ?", productID).Updates(updateData).Error
}

// PublishProductByShop publishes a product using sqlc
func (p *productRepository) PublishProductByShop(ctx context.Context, productID string, shopID string) error {
	_, err := p.sqlc.PublishProduct(ctx, database.PublishProductParams{
		ID:          productID,
		ProductShop: shopID,
	})
	return err
}

// UnPublishProductByShop unpublishes a product using sqlc
func (p *productRepository) UnPublishProductByShop(ctx context.Context, productID string, shopID string) error {
	_, err := p.sqlc.UnpublishProduct(ctx, database.UnpublishProductParams{
		ID:          productID,
		ProductShop: shopID,
	})
	return err
}

// FindAllDraftsForShop finds all draft products for a shop using sqlc
func (p *productRepository) FindAllDraftsForShop(ctx context.Context, shopID string, limit, offset int) ([]model.ProductModel, error) {
	drafts, err := p.sqlc.ListDraftProducts(ctx, database.ListDraftProductsParams{
		ProductShop: shopID,
		Limit:       int32(limit),
		Offset:      int32(offset),
	})
	if err != nil {
		return nil, err
	}
	
	var products []model.ProductModel
	for _, draft := range drafts {
		// Convert string price to float64
		price, _ := strconv.ParseFloat(draft.ProductPrice, 64)
		
		product := model.ProductModel{
			ID:                 draft.ID,
			ProductName:        draft.ProductName,
			ProductPrice:       price,
			ProductThumb:       draft.ProductThumb.String,
			ProductDescription: draft.ProductDescription.String,
			ProductQuantity:    int(draft.ProductQuantity),
			ProductType:        draft.ProductType,
			SubProductType:     draft.SubProductType.String,
			ProductStatus:      draft.ProductStatus,
			ProductSelled:      int(draft.ProductSelled),
			ProductShop:        draft.ProductShop,
			IsDraft:            draft.IsDraft,
			IsPublished:        draft.IsPublished,
		}
		
		// Handle nullables
		if draft.CreatedAt.Valid {
			product.CreatedAt = draft.CreatedAt.Time
		}
		
		if draft.UpdatedAt.Valid {
			product.UpdatedAt = draft.UpdatedAt.Time
		}
		
		// Handle discounted price
		if draft.ProductDiscountedPrice.Valid {
			discountedPrice, _ := strconv.ParseFloat(draft.ProductDiscountedPrice.String, 64)
			product.ProductDiscountPrice = discountedPrice
		}
		
		// Unmarshal JSON arrays
		if len(draft.ProductVideos) > 0 {
			var videos []string
			if err := json.Unmarshal(draft.ProductVideos, &videos); err == nil {
				product.ProductVideos = videos
			}
		}
		
		if len(draft.ProductPictures) > 0 {
			var pictures []string
			if err := json.Unmarshal(draft.ProductPictures, &pictures); err == nil {
				product.ProductPictures = pictures
			}
		}
		
		products = append(products, product)
	}
	
	return products, nil
}

// FindAllPublishForShop finds all published products for a shop using sqlc
func (p *productRepository) FindAllPublishForShop(ctx context.Context, shopID string, limit, offset int) ([]model.ProductModel, error) {
	published, err := p.sqlc.ListPublishedProducts(ctx, database.ListPublishedProductsParams{
		ProductShop: shopID,
		Limit:       int32(limit),
		Offset:      int32(offset),
	})
	if err != nil {
		return nil, err
	}
	
	var products []model.ProductModel
	for _, pub := range published {
		// Convert string price to float64
		price, _ := strconv.ParseFloat(pub.ProductPrice, 64)
		
		product := model.ProductModel{
			ID:                 pub.ID,
			ProductName:        pub.ProductName,
			ProductPrice:       price,
			ProductThumb:       pub.ProductThumb.String,
			ProductDescription: pub.ProductDescription.String,
			ProductQuantity:    int(pub.ProductQuantity),
			ProductType:        pub.ProductType,
			SubProductType:     pub.SubProductType.String,
			ProductStatus:      pub.ProductStatus,
			ProductSelled:      int(pub.ProductSelled),
			ProductShop:        pub.ProductShop,
			IsDraft:            pub.IsDraft,
			IsPublished:        pub.IsPublished,
		}
		
		// Handle nullables
		if pub.CreatedAt.Valid {
			product.CreatedAt = pub.CreatedAt.Time
		}
		
		if pub.UpdatedAt.Valid {
			product.UpdatedAt = pub.UpdatedAt.Time
		}
		
		// Handle discounted price
		if pub.ProductDiscountedPrice.Valid {
			discountedPrice, _ := strconv.ParseFloat(pub.ProductDiscountedPrice.String, 64)
			product.ProductDiscountPrice = discountedPrice
		}
		
		// Unmarshal JSON arrays
		if len(pub.ProductVideos) > 0 {
			var videos []string
			if err := json.Unmarshal(pub.ProductVideos, &videos); err == nil {
				product.ProductVideos = videos
			}
		}
		
		if len(pub.ProductPictures) > 0 {
			var pictures []string
			if err := json.Unmarshal(pub.ProductPictures, &pictures); err == nil {
				product.ProductPictures = pictures
			}
		}
		
		products = append(products, product)
	}
	
	return products, nil
}

// FindAllProducts finds all products based on params using sqlc and gorm
func (p *productRepository) FindAllProducts(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error) {
	// Set default pagination values
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}
	offset := (params.Page - 1) * params.Limit
	
	var products []model.ProductModel
	var totalCount int64
	
	// If there's a product type filter, use specific query
	if params.ProductType != "" {
		// Get count
		count, err := p.sqlc.CountProductsByType(ctx, params.ProductType)
		if err != nil {
			return nil, err
		}
		totalCount = count
		
		// Get products
		dbProducts, err := p.sqlc.ListProductsByType(ctx, database.ListProductsByTypeParams{
			ProductType: params.ProductType,
			Limit:       int32(params.Limit),
			Offset:      int32(offset),
		})
		if err != nil {
			return nil, err
		}
		
		// Convert to model
		for _, prod := range dbProducts {
			product := convertDbProductToModel(prod)
			products = append(products, product)
		}
	} else if params.Keyword != "" {
		// If there's a keyword search
		searchPattern := "%" + params.Keyword + "%"
		
		// Get count
		count, err := p.sqlc.CountSearchProductsByName(ctx, searchPattern)
		if err != nil {
			return nil, err
		}
		totalCount = count
		
		// Get products
		dbProducts, err := p.sqlc.SearchProductsByName(ctx, database.SearchProductsByNameParams{
			ProductName: searchPattern,
			Limit:       int32(params.Limit),
			Offset:      int32(offset),
		})
		if err != nil {
			return nil, err
		}
		
		// Convert to model
		for _, prod := range dbProducts {
			product := convertDbProductToModel(prod)
			products = append(products, product)
		}
	} else {
		// Default list all published products
		count, err := p.sqlc.CountAllPublishedProducts(ctx)
		if err != nil {
			return nil, err
		}
		totalCount = count
		
		dbProducts, err := p.sqlc.ListAllPublishedProducts(ctx, database.ListAllPublishedProductsParams{
			Limit:  int32(params.Limit),
			Offset: int32(offset),
		})
		if err != nil {
			return nil, err
		}
		
		// Convert to model
		for _, prod := range dbProducts {
			product := convertDbProductToModel(prod)
			products = append(products, product)
		}
	}
	
	totalPages := int(math.Ceil(float64(totalCount) / float64(params.Limit)))
	
	return &model.ProductResponse{
		CurrentPage: params.Page,
		TotalPages:  totalPages,
		Total:       int(totalCount),
		Data:        products,
	}, nil
}

// FindProductsByDiscount finds products ordered by discount
func (p *productRepository) FindProductsByDiscount(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error) {
	// Set default pagination values
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}
	offset := (params.Page - 1) * params.Limit
	
	// Get count of all published products
	count, err := p.sqlc.CountAllPublishedProducts(ctx)
	if err != nil {
		return nil, err
	}
	
	// Get products ordered by discount price
	dbProducts, err := p.sqlc.ListProductsByDiscount(ctx, database.ListProductsByDiscountParams{
		Limit:  int32(params.Limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	
	// Convert to model
	var products []model.ProductModel
	for _, prod := range dbProducts {
		product := convertDbProductToModel(prod)
		products = append(products, product)
	}
	
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))
	
	return &model.ProductResponse{
		CurrentPage: params.Page,
		TotalPages:  totalPages,
		Total:       int(count),
		Data:        products,
	}, nil
}

// FindProductsBySelled finds products ordered by number sold
func (p *productRepository) FindProductsBySelled(ctx context.Context, params *model.ProductQueryParams) (*model.ProductResponse, error) {
	// Set default pagination values
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}
	offset := (params.Page - 1) * params.Limit
	
	// Get count of all published products
	count, err := p.sqlc.CountAllPublishedProducts(ctx)
	if err != nil {
		return nil, err
	}
	
	// Get products ordered by selled count
	dbProducts, err := p.sqlc.ListProductsBySelled(ctx, database.ListProductsBySelledParams{
		Limit:  int32(params.Limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	
	// Convert to model
	var products []model.ProductModel
	for _, prod := range dbProducts {
		product := convertDbProductToModel(prod)
		products = append(products, product)
	}
	
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))
	
	return &model.ProductResponse{
		CurrentPage: params.Page,
		TotalPages:  totalPages,
		Total:       int(count),
		Data:        products,
	}, nil
}

// SearchProducts searches products by keyword using sqlc
func (p *productRepository) SearchProducts(ctx context.Context, keyword string) ([]model.ProductModel, error) {
	searchPattern := "%" + keyword + "%"
	
	dbProducts, err := p.sqlc.SearchProductsByName(ctx, database.SearchProductsByNameParams{
		ProductName: searchPattern,
		Limit:       100, // Reasonable limit
		Offset:      0,
	})
	if err != nil {
		return nil, err
	}
	
	// Convert to model
	var products []model.ProductModel
	for _, prod := range dbProducts {
		product := convertDbProductToModel(prod)
		products = append(products, product)
	}
	
	return products, nil
}

// CountProducts counts products based on query using gorm
func (p *productRepository) CountProducts(ctx context.Context, query map[string]interface{}) (int64, error) {
	var count int64
	err := p.db.Model(&model.ProductModel{}).Where(query).Count(&count).Error
	return count, err
}

// Helper function to convert database product to model
func convertDbProductToModel(dbProduct database.Product) model.ProductModel {
	// Convert string price to float64
	price, _ := strconv.ParseFloat(dbProduct.ProductPrice, 64)
	
	product := model.ProductModel{
		ID:                 dbProduct.ID,
		ProductName:        dbProduct.ProductName,
		ProductPrice:       price,
		ProductThumb:       dbProduct.ProductThumb.String,
		ProductDescription: dbProduct.ProductDescription.String,
		ProductQuantity:    int(dbProduct.ProductQuantity),
		ProductType:        dbProduct.ProductType,
		SubProductType:     dbProduct.SubProductType.String,
		ProductStatus:      dbProduct.ProductStatus,
		ProductSelled:      int(dbProduct.ProductSelled),
		ProductShop:        dbProduct.ProductShop,
		IsDraft:            dbProduct.IsDraft,
		IsPublished:        dbProduct.IsPublished,
	}
	
	// Handle nullables
	if dbProduct.CreatedAt.Valid {
		product.CreatedAt = dbProduct.CreatedAt.Time
	}
	
	if dbProduct.UpdatedAt.Valid {
		product.UpdatedAt = dbProduct.UpdatedAt.Time
	}
	
	// Handle discounted price
	if dbProduct.ProductDiscountedPrice.Valid {
		discountedPrice, _ := strconv.ParseFloat(dbProduct.ProductDiscountedPrice.String, 64)
		product.ProductDiscountPrice = discountedPrice
	}
	
	// Unmarshal JSON arrays
	if len(dbProduct.ProductVideos) > 0 {
		var videos []string
		if err := json.Unmarshal(dbProduct.ProductVideos, &videos); err == nil {
			product.ProductVideos = videos
		}
	}
	
	if len(dbProduct.ProductPictures) > 0 {
		var pictures []string
		if err := json.Unmarshal(dbProduct.ProductPictures, &pictures); err == nil {
			product.ProductPictures = pictures
		}
	}
	
	return product
} 
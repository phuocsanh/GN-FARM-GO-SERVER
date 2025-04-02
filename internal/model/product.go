package model

import "time"

// ProductModel là cấu trúc chung cho tất cả các sản phẩm
type ProductModel struct {
	ID                   string    `json:"id" gorm:"primaryKey"`
	ProductName          string    `json:"product_name"`
	ProductPrice         float64   `json:"product_price"`
	ProductDiscountPrice float64   `json:"product_discounted_price"`
	ProductThumb         string    `json:"product_thumb"`
	ProductDescription   string    `json:"product_description"`
	ProductQuantity      int       `json:"product_quantity"`
	ProductType          string    `json:"product_type"`
	SubProductType       string    `json:"sub_product_type"`
	ProductVideos        []string  `json:"product_videos"`
	ProductPictures      []string  `json:"product_pictures"`
	ProductStatus        string    `json:"product_status"`
	ProductSelled        int       `json:"product_selled"`
	ProductShop          string    `json:"product_shop"`
	IsDraft              bool      `json:"is_draft" gorm:"default:true"`
	IsPublished          bool      `json:"is_published" gorm:"default:false"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// TableName ghi đè tên bảng trong gorm
func (ProductModel) TableName() string {
	return "products"
}

// MushroomModel là mô hình cho sản phẩm loại nấm
type MushroomModel struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	ProductShop string  `json:"product_shop"`
	Weight      float64 `json:"weight"`
	Origin      string  `json:"origin"`
	Freshness   string  `json:"freshness"`
	PackageType string  `json:"package_type"`
}

// TableName ghi đè tên bảng trong gorm
func (MushroomModel) TableName() string {
	return "mushrooms"
}

// VegetableModel là mô hình cho sản phẩm loại rau củ
type VegetableModel struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	ProductShop string  `json:"product_shop"`
	Weight      float64 `json:"weight"`
	Origin      string  `json:"origin"`
	Freshness   string  `json:"freshness"`
	PackageType string  `json:"package_type"`
}

// TableName ghi đè tên bảng trong gorm
func (VegetableModel) TableName() string {
	return "vegetables"
}

// BonsaiModel là mô hình cho sản phẩm loại bonsai
type BonsaiModel struct {
	ID          string `json:"id" gorm:"primaryKey"`
	ProductShop string `json:"product_shop"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Style       string `json:"style"`
	Species     string `json:"species"`
	PotType     string `json:"pot_type"`
}

// TableName ghi đè tên bảng trong gorm
func (BonsaiModel) TableName() string {
	return "bonsais"
}

// ProductInput là cấu trúc cho dữ liệu đầu vào khi tạo sản phẩm
type ProductInput struct {
	ProductName          string                 `json:"product_name"`
	ProductPrice         float64                `json:"product_price"`
	ProductDiscountPrice float64                `json:"product_discounted_price"`
	ProductThumb         string                 `json:"product_thumb"`
	ProductDescription   string                 `json:"product_description"`
	ProductQuantity      int                    `json:"product_quantity"`
	ProductType          string                 `json:"product_type"`
	SubProductType       string                 `json:"sub_product_type"`
	ProductVideos        []string               `json:"product_videos"`
	ProductPictures      []string               `json:"product_pictures"`
	ProductStatus        string                 `json:"product_status"`
	ProductAttributes    map[string]interface{} `json:"product_attributes"`
}

// InventoryInput là cấu trúc cho dữ liệu đầu vào khi tạo inventory
type InventoryInput struct {
	ProductID string `json:"product_id"`
	ShopID    string `json:"shop_id"`
	Location  string `json:"location"`
	Stock     int    `json:"stock"`
}

// ProductQueryParams là cấu trúc cho tham số truy vấn sản phẩm
type ProductQueryParams struct {
	Page        int    `form:"page" json:"page"`
	Limit       int    `form:"limit" json:"limit"`
	ProductType string `form:"product_type" json:"product_type"`
	Keyword     string `form:"keyword" json:"keyword"`
}

// ProductResponse là cấu trúc dữ liệu phản hồi
type ProductResponse struct {
	CurrentPage int            `json:"current_page"`
	TotalPages  int            `json:"total_pages"`
	Total       int            `json:"total"`
	Data        []ProductModel `json:"data"`
} 
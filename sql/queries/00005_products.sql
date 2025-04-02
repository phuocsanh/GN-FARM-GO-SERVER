-- name: CreateProduct :execresult
INSERT INTO products (
    id, product_name, product_price, product_discounted_price, 
    product_thumb, product_description, product_quantity, 
    product_type, sub_product_type, product_videos, 
    product_pictures, product_status, product_shop, 
    is_draft, is_published
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateMushroom :execresult
INSERT INTO mushrooms (
    id, product_shop, weight, origin, freshness, package_type
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: CreateVegetable :execresult
INSERT INTO vegetables (
    id, product_shop, weight, origin, freshness, package_type
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: CreateBonsai :execresult
INSERT INTO bonsais (
    id, product_shop, age, height, style, species, pot_type
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateInventory :execresult
INSERT INTO inventory (
    product_id, shop_id, location, stock
) VALUES (
    ?, ?, ?, ?
);

-- name: GetProductByID :one
SELECT * FROM products
WHERE id = ? LIMIT 1;

-- name: GetMushroomByID :one
SELECT * FROM mushrooms
WHERE id = ? LIMIT 1;

-- name: GetVegetableByID :one
SELECT * FROM vegetables
WHERE id = ? LIMIT 1;

-- name: GetBonsaiByID :one
SELECT * FROM bonsais
WHERE id = ? LIMIT 1;

-- name: UpdateProduct :execresult
UPDATE products
SET 
    product_name = COALESCE(NULLIF(?, ''), product_name),
    product_price = COALESCE(NULLIF(?, 0), product_price),
    product_discounted_price = COALESCE(?, product_discounted_price),
    product_thumb = COALESCE(NULLIF(?, ''), product_thumb),
    product_description = COALESCE(NULLIF(?, ''), product_description),
    product_quantity = COALESCE(NULLIF(?, 0), product_quantity),
    sub_product_type = COALESCE(NULLIF(?, ''), sub_product_type),
    product_videos = COALESCE(?, product_videos),
    product_pictures = COALESCE(?, product_pictures),
    product_status = COALESCE(NULLIF(?, ''), product_status)
WHERE id = ? AND product_shop = ?;

-- name: UpdateMushroom :execresult
UPDATE mushrooms
SET 
    weight = COALESCE(?, weight),
    origin = COALESCE(NULLIF(?, ''), origin),
    freshness = COALESCE(NULLIF(?, ''), freshness),
    package_type = COALESCE(NULLIF(?, ''), package_type)
WHERE id = ? AND product_shop = ?;

-- name: UpdateVegetable :execresult
UPDATE vegetables
SET 
    weight = COALESCE(?, weight),
    origin = COALESCE(NULLIF(?, ''), origin),
    freshness = COALESCE(NULLIF(?, ''), freshness),
    package_type = COALESCE(NULLIF(?, ''), package_type)
WHERE id = ? AND product_shop = ?;

-- name: UpdateBonsai :execresult
UPDATE bonsais
SET 
    age = COALESCE(?, age),
    height = COALESCE(?, height),
    style = COALESCE(NULLIF(?, ''), style),
    species = COALESCE(NULLIF(?, ''), species),
    pot_type = COALESCE(NULLIF(?, ''), pot_type)
WHERE id = ? AND product_shop = ?;

-- name: PublishProduct :execresult
UPDATE products
SET is_draft = false, is_published = true
WHERE id = ? AND product_shop = ?;

-- name: UnpublishProduct :execresult
UPDATE products
SET is_draft = true, is_published = false
WHERE id = ? AND product_shop = ?;

-- name: ListDraftProducts :many
SELECT * FROM products
WHERE product_shop = ? AND is_draft = true
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListPublishedProducts :many
SELECT * FROM products
WHERE product_shop = ? AND is_published = true
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListAllPublishedProducts :many
SELECT * FROM products
WHERE is_published = true
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountAllPublishedProducts :one
SELECT COUNT(*) FROM products
WHERE is_published = true;

-- name: ListProductsByType :many
SELECT * FROM products
WHERE product_type = ? AND is_published = true
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountProductsByType :one
SELECT COUNT(*) FROM products
WHERE product_type = ? AND is_published = true;

-- name: SearchProductsByName :many
SELECT * FROM products
WHERE product_name LIKE ? AND is_published = true
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountSearchProductsByName :one
SELECT COUNT(*) FROM products
WHERE product_name LIKE ? AND is_published = true;

-- name: ListProductsByDiscount :many
SELECT * FROM products
WHERE is_published = true
ORDER BY product_discounted_price DESC
LIMIT ? OFFSET ?;

-- name: ListProductsBySelled :many
SELECT * FROM products
WHERE is_published = true
ORDER BY product_selled DESC
LIMIT ? OFFSET ?; 
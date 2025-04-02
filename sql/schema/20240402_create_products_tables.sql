-- +goose Up
-- +goose StatementBegin
-- Product table
CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) PRIMARY KEY,                    -- Product ID (UUID)
    product_name VARCHAR(255) NOT NULL,            -- Product name
    product_price DECIMAL(10, 2) NOT NULL,         -- Product price
    product_discounted_price DECIMAL(10, 2) NULL,  -- Discounted price
    product_thumb VARCHAR(255) NULL,               -- Product thumbnail URL
    product_description TEXT NULL,                 -- Product description
    product_quantity INT NOT NULL DEFAULT 0,       -- Quantity in stock
    product_type VARCHAR(50) NOT NULL,             -- Product type (Mushroom, Vegetable, Bonsai)
    sub_product_type VARCHAR(50) NULL,             -- Sub-category
    product_videos JSON NULL,                      -- Array of video URLs
    product_pictures JSON NULL,                    -- Array of picture URLs
    product_status VARCHAR(20) NOT NULL DEFAULT 'active', -- Product status
    product_selled INT NOT NULL DEFAULT 0,         -- Number of items sold
    product_shop VARCHAR(36) NOT NULL,             -- Shop ID (User ID)
    is_draft BOOLEAN NOT NULL DEFAULT TRUE,        -- Whether it's a draft
    is_published BOOLEAN NOT NULL DEFAULT FALSE,   -- Whether it's published
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Update time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Products table';

-- Mushroom products
CREATE TABLE IF NOT EXISTS mushrooms (
    id VARCHAR(36) PRIMARY KEY,                -- Product ID (UUID)
    product_shop VARCHAR(36) NOT NULL,         -- Shop ID (User ID)
    weight DECIMAL(10, 2) NULL,                -- Weight in grams/kg
    origin VARCHAR(100) NULL,                  -- Origin location
    freshness VARCHAR(50) NULL,                -- Freshness description
    package_type VARCHAR(50) NULL,             -- Type of packaging
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    FOREIGN KEY (id) REFERENCES products(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Mushroom products table';

-- Vegetable products
CREATE TABLE IF NOT EXISTS vegetables (
    id VARCHAR(36) PRIMARY KEY,                -- Product ID (UUID)
    product_shop VARCHAR(36) NOT NULL,         -- Shop ID (User ID)
    weight DECIMAL(10, 2) NULL,                -- Weight in grams/kg
    origin VARCHAR(100) NULL,                  -- Origin location
    freshness VARCHAR(50) NULL,                -- Freshness description
    package_type VARCHAR(50) NULL,             -- Type of packaging
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    FOREIGN KEY (id) REFERENCES products(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Vegetable products table';

-- Bonsai products
CREATE TABLE IF NOT EXISTS bonsais (
    id VARCHAR(36) PRIMARY KEY,                -- Product ID (UUID)
    product_shop VARCHAR(36) NOT NULL,         -- Shop ID (User ID)
    age INT NULL,                              -- Age in years
    height INT NULL,                           -- Height in cm
    style VARCHAR(50) NULL,                    -- Bonsai style
    species VARCHAR(100) NULL,                 -- Plant species
    pot_type VARCHAR(50) NULL,                 -- Type of pot
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    FOREIGN KEY (id) REFERENCES products(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Bonsai products table';

-- Inventory table
CREATE TABLE IF NOT EXISTS inventory (
    id INT AUTO_INCREMENT PRIMARY KEY,         -- Inventory ID
    product_id VARCHAR(36) NOT NULL,           -- Product ID (UUID)
    shop_id VARCHAR(36) NOT NULL,              -- Shop ID (User ID)
    location VARCHAR(255) NULL,                -- Storage location
    stock INT NOT NULL DEFAULT 0,              -- Current stock
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Inventory table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `inventory`;
DROP TABLE IF EXISTS `bonsais`;
DROP TABLE IF EXISTS `vegetables`;
DROP TABLE IF EXISTS `mushrooms`;
DROP TABLE IF EXISTS `products`;
-- +goose StatementEnd 
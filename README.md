# Base Project - Go Clean Architecture

Đây là một base project được xây dựng theo Clean Architecture pattern, sử dụng Gin framework và GORM cho việc phát triển REST API.

## Cấu trúc Project

```
├── common/                 # Shared utilities và common types
├── component/             # External components (JWT, etc.)
├── middleware/            # HTTP middleware
├── module/               # Business modules
│   ├── user/            # User module (authentication, authorization)
│   └── upload/          # File upload module
├── static/              # Static files
└── main.go             # Application entry point
```

## Tính năng có sẵn

### Authentication & Authorization
- User registration
- User login với JWT token
- Protected routes với middleware
- Role-based access control (User, Admin, Mod, Shipper)

### File Upload
- Upload files với endpoint `/v1/upload`

### Database
- MySQL integration với GORM
- Auto-migration support
- Common SQL model với masking

## Cài đặt và Chạy

### Yêu cầu
- Go 1.24+
- MySQL database (hoặc Docker)

### Cách 1: Sử dụng Docker (Khuyến nghị)

```bash
# Khởi động database và Redis
docker-compose up -d

# Cài đặt dependencies
go mod tidy

# Chạy ứng dụng
go run main.go
```

#### Lệnh Docker Compose hữu ích

```bash
# Build images
docker compose build

# Build và chạy nền
docker compose up -d --build

# Rebuild chỉ service app
docker compose build app && docker compose up -d app

# Xem logs của app
docker compose logs -f app
```

### Cách 2: Cài đặt thủ công

1. Cài đặt MySQL và tạo database
2. Tạo file `.env` từ `env.example` và cấu hình:

```bash
DB_CONN=app_user:app_password@tcp(localhost:3306)/base_project?charset=utf8mb4&parseTime=True&loc=Local
SECRET=your_super_secret_jwt_key_here
```

3. Chạy ứng dụng:

```bash
# Cài đặt dependencies
go mod tidy

# Chạy ứng dụng
go run main.go
```

Ứng dụng sẽ chạy trên port 3000.

### Development Tools

```bash
# Cài đặt development tools
make install-tools

# Chạy với hot reload
make dev

# Build và run
make run

# Chạy tests
make test
```

## API Endpoints

### Authentication
- `POST /v1/register` - Đăng ký user mới
- `POST /v1/login` - Đăng nhập
- `GET /v1/profile` - Lấy thông tin profile (cần authentication)

### Upload
- `PUT /v1/upload` - Upload file

### Category (CRUD)
- `POST /v1/categories` - Tạo category
  - Body ví dụ:
    ```json
    {
      "category_name": "Coffee Beans",
      "description": "All kinds of beans",
      "status": "active"
    }
    ```
- `GET /v1/categories` - Danh sách category (paging `page`, `limit`)
- `GET /v1/categories/:id` - Lấy chi tiết category
- `PUT /v1/categories/:id` - Cập nhật category (partial)
  - Body ví dụ:
    ```json
    {
      "description": "Premium roasted beans"
    }
    ```
- `DELETE /v1/categories/:id` - Deactive category (soft delete, đặt `status = "deactive"`)

### Utility
- `GET /ping` - Health check

## Phát triển Module mới

### 1. Tạo cấu trúc module
```
module/
└── your_module/
    ├── biz/           # Business logic
    ├── model/         # Data models
    ├── storage/       # Data access layer
    └── transport/     # HTTP handlers
```

### 2. Ví dụ tạo module Product

**Model** (`module/product/model/product.go`):
```go
package model

import "base-project/common"

type Product struct {
    common.SQLModel
    Name        string `json:"name" gorm:"column:name;"`
    Description string `json:"description" gorm:"column:description;"`
    Price       int    `json:"price" gorm:"column:price;"`
}

func (Product) TableName() string {
    return "products"
}
```

**Storage** (`module/product/storage/store.go`):
```go
package storage

import (
    "base-project/common"
    "base-project/module/product/model"
    "context"
    "gorm.io/gorm"
)

type sqlStore struct {
    db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
    return &sqlStore{db: db}
}

func (s *sqlStore) CreateProduct(ctx context.Context, data *model.Product) error {
    return s.db.Create(data).Error
}
```

**Handler** (`module/product/transport/gin/product_handler.go`):
```go
package gin

import (
    "base-project/common"
    "base-project/module/product/biz"
    "base-project/module/product/model"
    "base-project/module/product/storage"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var data model.Product
        
        if err := c.ShouldBind(&data); err != nil {
            c.JSON(http.StatusBadRequest, common.NewErrorResponse(err))
            return
        }
        
        store := storage.NewSQLStore(db)
        biz := biz.NewCreateProductBiz(store)
        
        if err := biz.CreateProduct(c.Request.Context(), &data); err != nil {
            c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
            return
        }
        
        c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
    }
}
```

### 3. Thêm routes vào main.go

```go
// Import module mới
productGin "base-project/module/product/transport/gin"

// Trong v1 group
products := v1.Group("/products", midAuth)
{
    products.POST("", productGin.CreateProduct(db))
    products.GET("", productGin.ListProducts(db))
    products.GET("/:id", productGin.GetProduct(db))
    products.PATCH("/:id", productGin.UpdateProduct(db))
    products.DELETE("/:id", productGin.DeleteProduct(db))
}
```

## Best Practices

1. **Clean Architecture**: Tuân thủ dependency rule, business logic không phụ thuộc vào framework
2. **Error Handling**: Sử dụng custom error types trong `common/app_err.go`
3. **Response Format**: Sử dụng common response format trong `common/app_response.go`
4. **Database**: Sử dụng GORM với auto-migration
5. **Authentication**: JWT-based với middleware protection
6. **File Upload**: Sử dụng module upload có sẵn

## License

MIT License
# Mocau Backend API

Go REST API với Clean Architecture, sử dụng Gin framework và GORM.

## Cấu trúc Project

```
├── common/                 # Shared utilities
├── component/             # External components (JWT)
├── middleware/            # HTTP middleware
├── module/               # Business modules
│   ├── user/            # User module
│   ├── category/        # Category module
│   ├── product/         # Product module
│   ├── blog/            # Blog module
│   ├── order/           # Order module
│   └── upload/          # File upload module
├── static/              # Static files
└── main.go             # Application entry point
```

## Tính năng

- **Authentication**: JWT-based với middleware protection
- **User Management**: Registration, login, profile
- **Category CRUD**: Create, read, update, delete categories
- **Product CRUD**: Create, read, update products
- **Blog CRUD**: Create, read, update, delete blogs
- **Order Management**: Simple order system (customer info → order → staff confirmation)
- **File Upload**: Upload files với endpoint `/v1/upload`
- **Swagger Documentation**: Auto-generated API docs
- **Database**: MySQL với GORM auto-migration

## Cài đặt và Chạy

### Yêu cầu
- Go 1.23+
- MySQL database

### Cách 1: Docker Compose (Khuyến nghị)

```bash
# Tạo file .env từ template
cp env.example .env

# Chỉnh sửa .env với thông tin database
nano .env

# Build và chạy
docker compose up -d --build
```

### Cách 2: Chạy local

```bash
# Tạo file .env
cp env.example .env

# Chỉnh sửa .env với thông tin MySQL
nano .env

# Cài dependencies
go mod tidy

# Chạy ứng dụng
go run main.go
```

## API Endpoints

### Authentication
- `POST /v1/register` - Đăng ký user
- `POST /v1/login` - Đăng nhập
- `GET /v1/profile` - Lấy profile (cần auth)

### Category CRUD
- `POST /v1/categories` - Tạo category
- `GET /v1/categories` - Danh sách category (paging)
- `GET /v1/categories/:id` - Lấy chi tiết category
- `PUT /v1/categories/:id` - Cập nhật category
- `DELETE /v1/categories/:id` - Xóa category (soft delete)

### Upload
- `PUT /v1/upload` - Upload file

### Documentation
- `GET /swagger/index.html` - Swagger UI

## Environment Variables

```bash
# Database
DB_CONN=username:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local

# JWT Secret
SECRET=your_secret_key

# Swagger Host
SWAGGER_HOST=localhost:3000

# Server
PORT=3000
GIN_MODE=debug
```

## Swagger Documentation

```bash
# Cài swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g main.go -o ./docs

# Mở Swagger UI
# http://localhost:3000/swagger/index.html
```

## Development

### Tạo module mới

1. Tạo cấu trúc thư mục:
```
module/your_module/
├── biz/           # Business logic
├── model/         # Data models
├── storage/       # Data access
└── transport/     # HTTP handlers
```

2. Thêm routes vào `main.go`

### Database Migration

App tự động migrate tables khi khởi động. Để tạo schema thủ công:

```sql
-- Xem file init.sql
```

## License

MIT License
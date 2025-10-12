package main

import (
	"log"
	"mocau-backend/component/tokenprovider/jwt"
	"mocau-backend/docs"
	"mocau-backend/middleware"
	blogModel "mocau-backend/module/blog/model"
	blogGin "mocau-backend/module/blog/transport/ginBlog"
	catModel "mocau-backend/module/category/model"
	catGin "mocau-backend/module/category/transport/ginCategory"
	orderModel "mocau-backend/module/order/model"
	orderGin "mocau-backend/module/order/transport/ginOrder"
	orderitemModel "mocau-backend/module/orderitem/model"
	orderitemGin "mocau-backend/module/orderitem/transport/ginOrderItem"
	prodModel "mocau-backend/module/product/model"
	prodGin "mocau-backend/module/product/transport/ginProduct"
	"mocau-backend/module/upload"
	userModel "mocau-backend/module/user/model"
	userBiz "mocau-backend/module/user/biz"
	"mocau-backend/module/user/storage"
	"mocau-backend/module/user/transport/ginUser"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Mocau Backend API
// @version 1.0
// @description API documentation for Mocau Backend
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Swagger docs
	docs.SwaggerInfo.Title = "Mocau Backend API"
	docs.SwaggerInfo.Description = "API documentation for Mocau Backend"
	docs.SwaggerInfo.Version = "1.0"
	// Set Swagger host based on environment
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:3000" // Default for local development
	}
	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	dsn := os.Getenv("DB_CONN")
	systemSecret := os.Getenv("SECRET")

	// Debug: Check if environment variables are loaded
	if dsn == "" {
		log.Fatal("DB_CONN environment variable is not set. Please check your .env file or environment variables.")
	}
	if systemSecret == "" {
		log.Fatal("SECRET environment variable is not set. Please check your .env file or environment variables.")
	}

	log.Printf("Connecting to database with DSN: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	log.Println("Connected to database", db)

	// Auto migrate database tables
	err = db.AutoMigrate(
		&userModel.User{},
		&catModel.Category{},
		&prodModel.Product{},
		&blogModel.Blog{},
		&orderModel.Order{},
		&orderitemModel.OrderItem{},
	)
	if err != nil {
		log.Fatalln("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")

	///////
	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenJwtProvider("jwt", systemSecret)
	// midAuth := middleware.RequiredAuth(authStore, tokenProvider) // Uncomment when needed

	// Initialize user business logic
	userListBiz := userBiz.NewListUserBusiness(authStore)
	userUpdateStatusBiz := userBiz.NewUpdateUserStatusBusiness(authStore)
	userUpdateProfileBiz := userBiz.NewUpdateUserProfileBusiness(authStore)

	r := gin.Default()
	r.Use(middleware.Recover())

	// Add CORS middleware - Full CORS support
	r.Use(func(c *gin.Context) {
		// Allow all origins
		c.Header("Access-Control-Allow-Origin", "*")
		
		// Allow all methods
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
		
		// Allow all headers
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		
		// Allow credentials
		c.Header("Access-Control-Allow-Credentials", "true")
		
		// Expose headers
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		
		// Cache preflight for 12 hours
		c.Header("Access-Control-Max-Age", "43200")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Serve static files (backward compatibility)
	r.Static("/static", "./static")
	// Serve media files from VPS /media directory
	r.Static("/media", "/media")

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db))

		v1.POST("/register", ginUser.Register(db))
		v1.POST("/login", ginUser.Login(db, tokenProvider))
		v1.GET("/profile", middleware.RequiredAuth(authStore, tokenProvider), ginUser.Profile())
		v1.GET("/users", middleware.RequiredAuth(authStore, tokenProvider), ginUser.ListUsers(userListBiz))
		v1.PUT("/users/:id/toggle-status", middleware.RequiredAuth(authStore, tokenProvider), ginUser.ToggleUserStatus(userUpdateStatusBiz))
		v1.PUT("/users/:id/profile", middleware.RequiredAuth(authStore, tokenProvider), ginUser.UpdateProfile(userUpdateProfileBiz))

		// Category routes
		v1.POST("/categories", middleware.RequiredAuth(authStore, tokenProvider), catGin.CreateCategory(db))
		v1.GET("/categories", catGin.ListCategories(db))
		v1.GET("/categories/:id", catGin.GetCategory(db))
		v1.PUT("/categories/:id", middleware.RequiredAuth(authStore, tokenProvider), catGin.UpdateCategory(db))
		v1.DELETE("/categories/:id", middleware.RequiredAuth(authStore, tokenProvider), catGin.DeleteCategory(db))

        // Product routes
        v1.POST("/products", middleware.RequiredAuth(authStore, tokenProvider), prodGin.CreateProduct(db))
        v1.GET("/products", prodGin.ListProducts(db))
        v1.GET("/products/top-selling", prodGin.GetTopSellingProducts(db))
        v1.GET("/products/revenue-growth", prodGin.GetRevenueGrowth(db))
        v1.GET("/products/:id", prodGin.GetProduct(db))
        v1.PUT("/products/:id", middleware.RequiredAuth(authStore, tokenProvider), prodGin.UpdateProduct(db))

        // Blog routes
        v1.POST("/blogs", middleware.RequiredAuth(authStore, tokenProvider), blogGin.CreateBlog(db))
        v1.GET("/blogs", blogGin.ListBlogs(db))
        v1.GET("/blogs/:id", blogGin.GetBlog(db))
        v1.PUT("/blogs/:id", middleware.RequiredAuth(authStore, tokenProvider), blogGin.UpdateBlog(db))
        v1.DELETE("/blogs/:id", middleware.RequiredAuth(authStore, tokenProvider), blogGin.DeleteBlog(db))

        // Order routes
        v1.POST("/orders", orderGin.CreateOrder(db))
        v1.GET("/orders", middleware.RequiredAuth(authStore, tokenProvider), orderGin.ListOrders(db))
        v1.GET("/orders/stats", middleware.RequiredAuth(authStore, tokenProvider), orderGin.GetOrderStats(db))
        v1.GET("/orders/search", middleware.RequiredAuth(authStore, tokenProvider), orderGin.SearchOrders(db))
        v1.GET("/orders/number/:order_number", orderGin.GetOrderByOrderNumber(db))
        v1.GET("/orders/:id", middleware.RequiredAuth(authStore, tokenProvider), orderGin.GetOrder(db))
        v1.PUT("/orders/:id", middleware.RequiredAuth(authStore, tokenProvider), orderGin.UpdateOrder(db))
        v1.PUT("/orders/:id/status", middleware.RequiredAuth(authStore, tokenProvider), orderGin.UpdateOrderStatus(db))
        v1.DELETE("/orders/:id", middleware.RequiredAuth(authStore, tokenProvider), orderGin.DeleteOrder(db))

        // Order Item routes
        v1.POST("/order-items", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.CreateOrderItem(db))
        v1.POST("/order-items/bulk", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.BulkCreateOrderItems(db))
        v1.GET("/order-items/order/:order_id", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.ListOrderItemsByOrder(db))
        v1.GET("/order-items/:id", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.GetOrderItem(db))
        v1.PUT("/order-items/:id", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.UpdateOrderItem(db))
        v1.PUT("/order-items/:id/quantity", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.UpdateOrderItemQuantity(db))
        v1.PUT("/order-items/bulk", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.BulkUpdateOrderItems(db))
        v1.DELETE("/order-items/:id", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.DeleteOrderItem(db))
        v1.DELETE("/order-items/bulk", middleware.RequiredAuth(authStore, tokenProvider), orderitemGin.BulkDeleteOrderItems(db))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}

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

	r := gin.Default()
	r.Use(middleware.Recover())

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

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

		// Category routes
		v1.POST("/categories", middleware.RequiredAuth(authStore, tokenProvider), catGin.CreateCategory(db))
		v1.GET("/categories", catGin.ListCategories(db))
		v1.GET("/categories/:id", catGin.GetCategory(db))
		v1.PUT("/categories/:id", middleware.RequiredAuth(authStore, tokenProvider), catGin.UpdateCategory(db))
		v1.DELETE("/categories/:id", middleware.RequiredAuth(authStore, tokenProvider), catGin.DeleteCategory(db))

        // Product routes
        v1.POST("/products", middleware.RequiredAuth(authStore, tokenProvider), prodGin.CreateProduct(db))
        v1.GET("/products", prodGin.ListProducts(db))
        v1.GET("/products/:id", prodGin.GetProduct(db))
        v1.PUT("/products/:id", middleware.RequiredAuth(authStore, tokenProvider), prodGin.UpdateProduct(db))

        // Blog routes
        v1.POST("/blogs", middleware.RequiredAuth(authStore, tokenProvider), blogGin.CreateBlog(db))
        v1.GET("/blogs", blogGin.ListBlogs(db))
        v1.GET("/blogs/:id", blogGin.GetBlog(db))
        v1.PUT("/blogs/:id", middleware.RequiredAuth(authStore, tokenProvider), blogGin.UpdateBlog(db))
        v1.DELETE("/blogs/:id", middleware.RequiredAuth(authStore, tokenProvider), blogGin.DeleteBlog(db))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}

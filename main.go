package main

import (
	"log"
	"mocau-backend/component/tokenprovider/jwt"
	"mocau-backend/docs"
	"mocau-backend/middleware"
	"mocau-backend/module/upload"
	"mocau-backend/module/user/model"
	"mocau-backend/module/user/storage"
	"mocau-backend/module/user/transport/ginUser"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	// Initialize Swagger docs
	docs.SwaggerInfo.Title = "Mocau Backend API"
	docs.SwaggerInfo.Description = "API documentation for Mocau Backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/v1"

	dsn := os.Getenv("DB_CONN")
	systemSecret := os.Getenv("SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to database", db)

	// Auto migrate database tables
	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatalln("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")

	///////
	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenJwtProvider("jwt", systemSecret)
	// midAuth := middleware.RequiredAuth(authStore, tokenProvider) // Uncomment when needed

	r := gin.Default()
	r.Use(middleware.Recover())

	r.Static("/static", "./static")

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db))

		v1.POST("/register", ginUser.Register(db))
		v1.POST("/login", ginUser.Login(db, tokenProvider))
		v1.GET("/profile", middleware.RequiredAuth(authStore, tokenProvider), ginUser.Profile())

		// TODO: Add your custom routes here
		// Example:
		// protected := v1.Group("/", midAuth)
		// {
		//     protected.GET("/your-endpoint", yourHandler)
		// }
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}

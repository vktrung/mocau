package ginUser

import (
	"mocau-backend/common"
	"mocau-backend/module/user/biz"
	"mocau-backend/module/user/model"
	"mocau-backend/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserCreate true "User registration data"
// @Success 200 {object} common.Response{data=int} "User created successfully"
// @Failure 400 {object} common.Response "Invalid request data"
// @Failure 409 {object} common.Response "Email or username already exists"
// @Router /register [post]
func Register(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data model.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()
		biz := biz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

    		// Return raw ID after successful creation
    		c.JSON(http.StatusOK, common.SimpleSuccessRes(data.Id))
	}
}

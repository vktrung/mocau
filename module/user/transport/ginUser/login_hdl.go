package ginUser

import (
	"mocau-backend/common"
	"mocau-backend/component/tokenprovider"
	"mocau-backend/module/user/biz"
	"mocau-backend/module/user/model"
	"mocau-backend/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login godoc
// @Summary User login
// @Description Login with username/email and password to get access token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body model.UserLogin true "Login credentials"
// @Success 200 {object} common.Response{data=tokenprovider.Token} "Login successful"
// @Failure 400 {object} common.Response "Invalid request data"
// @Failure 401 {object} common.Response "Invalid credentials"
// @Router /login [post]
func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		business := biz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(account))
	}
}

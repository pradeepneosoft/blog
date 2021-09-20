package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service    service.UserService
	jwtService service.JwtService
}

func NewUserController(s service.UserService, jwt service.JwtService) UserController {
	return UserController{
		service:    s,
		jwtService: jwt,
	}
}
func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.UserRegister
	if err := ctx.ShouldBind(&user); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Json")
		return
	}
	HashPassword, _ := util.HashPassword(user.Password)
	user.Password = HashPassword
	err := u.service.CreateUser(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "failed to cretae user")
		return
	}
	util.SuccessJSON(ctx, http.StatusOK, "successfully user created")

}

// func (u *UserController) Login(c *gin.Context) {
// 	var user models.UserLogin
// 	var hmacSampleSecret []byte
// 	if err := c.ShouldBind(&user); err != nil {
// 		util.ErrorJSON(c, http.StatusBadRequest, "invalid json")
// 		return
// 	}
// 	dbUser, err := u.service.LoginUser(user)
// 	if err != nil {
// 		util.ErrorJSON(c, http.StatusBadRequest, "INvalid Credential")
// 		return

// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user": dbUser,
// 		"exp":  time.Now().Add(time.Minute * 15).Unix(),
// 	})
// 	tokenString, err := token.SignedString(hmacSampleSecret)
// 	if err != nil {
// 		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
// 	}
// 	response := &util.Response{
// 		Success: true,
// 		Message: "Tokenm generated Successfully",
// 		Data:    tokenString,
// 	}
// 	c.JSON(http.StatusOK, response)

// }
func (u *UserController) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBind(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "invalid json")
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "INvalid Credential")
		return

	}
	tokenString, err := u.jwtService.GenerateToken(dbUser)
	fmt.Println("this", tokenString)
	if err != nil {
		fmt.Println("this err", err)
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Tokenm generated Successfully",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)

}
func (u *UserController) GetUsers(ctx *gin.Context) {
	var users models.User
	keyword := ctx.Query("keyword")
	data, total, err := u.service.FindAllUser(users, keyword)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find Questions")
		return
	}

	respArr := make([]map[string]interface{}, 0, 0)
	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)

	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "user result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

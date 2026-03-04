package controllers

import (
	"net/http"
	"todo-list/models"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// 创建用户
func (uc *UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 密码加密
	user.Password = utils.HashPassword(user.Password)

	if err := uc.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	user.Password = "" // 返回数据时隐藏密码
	utils.SuccessResponse(c, user)
}

// 用户登录
func (uc *UserController) Login(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := uc.DB.Where("username = ?", loginInfo.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if !utils.CheckPassword(loginInfo.Password, user.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 生成 token
	token := utils.GenerateToken(user.ID)
	utils.SuccessResponse(c, gin.H{"token": token})
}

// 获取用户信息
func (uc *UserController) GetInfo(c *gin.Context) {
	userID := c.GetUint("userID") // 从中间件获取用户ID
	var user models.User

	if err := uc.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "用户不存在")
		return
	}

	user.Password = "" // 隐藏密码
	utils.SuccessResponse(c, user)
}
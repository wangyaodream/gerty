package services

import (
	"fmt"
	"gerty/internal/dao"
	"gerty/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	//账号校验，检查账号是否存在
	accountDao := dao.NewAccountDao(c.db)
	fmt.Println("accountDao 创建完成")
	isExists, err := accountDao.IsExists(req.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if isExists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号已存在"})
		return
	}
	// 持久化
	if err := accountDao.Create(model.Account{
		UserID:   req.UserID,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Ct:       time.Now(),
		Ut:       time.Now(),
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &RegisterRsp{
			Message: fmt.Sprintln("success"),
		},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password error = %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

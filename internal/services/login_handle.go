package services

import (
	"gerty/internal/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		userID   = req.UserID
		password = req.Password
	)

	accountDao := dao.NewAccountDao(c.db)
	account, err := accountDao.FirstByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号不存在!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确!"})
		return
	}
	sessionId := generateSession()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LoginRsp{
			SessionID: sessionId,
			UserID:    account.UserID,
			Nickname:  account.Nickname,
		},
	})

	return
}

func generateSession() string {
	// TODO : session id 的生成
	// TODO : session id 持久化
	return "session-id"

}

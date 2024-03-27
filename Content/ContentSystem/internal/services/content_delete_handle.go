package services

import (
	"gerty/internal/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentDeleteReq struct {
	ID int `json:"id" binding:"required"` // content id
}

type ContentDeleteRsp struct {
	Message string `json:"message"`
}

func (c *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建DAO
	contentDao := dao.NewContentDao(c.db)
	ok, err := contentDao.IsExists(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "内容不存在"})
		return
	}

	if err := contentDao.Delete(req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentDeleteRsp{
			Message: "OK",
		},
	})
}

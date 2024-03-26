package services

import (
	"fmt"
	"gerty/internal/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Content struct {
	ID             int           `json:"id"`                        // 内容ID
	Title          string        `json:"title"`                     // 内容标题
	VideoURL       string        `json:"video_url" `                // 视频播放URL
	Author         string        `json:"author" binding:"required"` // 作者
	Description    string        `json:"description"`               // 内容描述
	Thumbnail      string        `json:"thumbnail"`                 // 封面图URL
	Category       string        `json:"category"`                  // 内容分类
	Duration       time.Duration `json:"duration"`                  // 内容时长
	Resolution     string        `json:"resolution"`                // 分辨率 如720p、1080p
	FileSize       int64         `json:"fileSize"`                  // 文件大小
	Format         string        `json:"format"`                    // 文件格式 如MP4、AVI
	Quality        int           `json:"quality"`                   // 视频质量 1-高清 2-标清
	ApprovalStatus int           `json:"approval_status"`
}

type ContentFindReq struct {
	ID       int `json:"id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type ContentFindRsp struct {
	Message  string    `json:"message"`
	Contents []Content `json:"contents"`
	Total    int64     `json:"total"`
}

func (c *CmsApp) ContentFind(ctx *gin.Context) {
	var req ContentFindReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 数据库交互
	contentDao := dao.NewContentDao(c.db)
	contentList, total, err := contentDao.Find(&dao.FindParams{
		ID:       req.ID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	contents := make([]Content, 0, len(contentList))

	// 向contents中填充实际数据
	for _, content := range contentList {
		contents = append(contents, Content{
			ID:             content.ID,
			Title:          content.Title,
			VideoURL:       content.VideoURL,
			Author:         content.Author,
			Description:    content.Description,
			Thumbnail:      content.Thumbnail,
			Category:       content.Category,
			Duration:       content.Duration,
			Resolution:     content.Resolution,
			FileSize:       content.FileSize,
			Format:         content.Format,
			Quality:        content.Quality,
			ApprovalStatus: content.ApprovalStatus,
		})
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentFindRsp{
			Message:  fmt.Sprintln("OK"),
			Contents: contents,
			Total:    total,
		},
	})
}

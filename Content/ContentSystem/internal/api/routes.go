package api

import (
	"gerty/internal/services"

	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api/"
	noAuthPath = "/out/api/"
)

func CmsRouter(r *gin.Engine) {

	cmsApp := services.NewCmsApp()
	// cmsApp := services.CmsApp{}
	// 创建一个中间件
	// session := &SessionAuth{}
	session := NewSessionAuth()

	// 创建路由组
	// 使用Use方法来注册中间件,在root下的所有接口都需要通过session.Auth
	root := r.Group(rootPath).Use(session.Auth)
	{
		// 运行逻辑绑定到一个特定的Hello方法中
		// root.GET("/cms/hello", cmsApp.Hello)
		root.GET("/cms/hello", cmsApp.Hello)
		// 内容生成
		root.POST("/cms/content/create", cmsApp.ContentCreate)
		root.POST("/cms/content/update", cmsApp.ContentUpdate)
		root.POST("/cms/content/delete", cmsApp.ContentDelete)
		root.POST("/cms/content/find", cmsApp.ContentFind)
	}

	noAuth := r.Group(noAuthPath)
	{
		noAuth.POST("/cms/register", cmsApp.Register)
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}

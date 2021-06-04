package routers

import (
	"blog/controller/api/v1"
	"github.com/gin-gonic/gin"
)

//注册分类路由
func RegisterCommentsRouter(router *gin.RouterGroup) {
	commentsRouter := router.Group("/comments")
	{
			//列表
			commentsRouter.GET("", v1.GetArticles)
			//分类ById
			commentsRouter.GET(":id", v1.GetArticle)
			//新建
			commentsRouter.POST("", v1.AddArticle)
			//更新ById
			commentsRouter.PUT(":id", v1.EditArticle)
			//删除ById
			commentsRouter.DELETE(":id", v1.DeleteArticle)

	}
}

package routers

import (
	v1 "blog/controller/api/v1"

	"github.com/gin-gonic/gin"
)

//注册分类路由
func RegisterCommentsRouter(router *gin.RouterGroup) {
	commentsRouter := router.Group("/comments")
	{
		//列表
		commentsRouter.GET("", v1.GetComments)
		//评论ById
		commentsRouter.GET(":id", v1.GetComment)
		//新建
		commentsRouter.POST("", v1.AddComment)
		//更新ById
		commentsRouter.PUT(":id", v1.EditComment)
		//删除ById
		commentsRouter.DELETE(":id", v1.DeleteComment)

	}
}

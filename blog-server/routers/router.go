package routers

import (
	"blog/middleware"
	"blog/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	v1 "blog/controller/api/v1"
	_ "blog/docs"
	"blog/middleware/jwt"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())      //日志
	r.Use(middleware.Cors()) // 跨域请求
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode) //设置运行模式

	//获取登录token
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //api注释文档
	apiv1 := r.Group("/api/v1")
	pub := apiv1.Group("/pub")
	{
		//获取验证码Id
		pub.GET("captchaid", v1.GetCaptcha)
		//获取验证码图片
		pub.GET("captcha", v1.ResCaptcha)
	}
	//登陆
	apiv1.POST("auth", v1.Auth)
	//注册
	apiv1.POST("reg", v1.Reg)
	apiv1.Use(jwt.JWT()) //令牌 验证中间件
	{
		//获取登录用户信息
		apiv1.GET("currentuser", v1.CurrentUser)
		//刷新token
		apiv1.GET("refreshtoken", v1.RefreshToken)
		//用户登出
		apiv1.POST("logout", v1.Logout)
		//登录用户修改密码
		apiv1.POST("password", v1.Password)
		//标签
		tag := apiv1.Group("/tags")
		{
			//列表
			tag.GET("", v1.GetTags)
			//新建
			tag.POST("", v1.AddTag)
			//更新
			tag.PUT(":id", v1.EditTag)
			//删除
			tag.DELETE(":id", v1.DeleteTag)
		}
		//分类
		category := apiv1.Group("/categories")
		{
			//列表
			category.GET("", v1.GetCategories)
			//新建
			category.POST("", v1.AddCategory)
			//更新
			category.PUT(":id", v1.EditCategory)
			//删除
			category.DELETE(":id", v1.DelCategory)
		}
		// 注册文章路由
		RegisterArticleRouter(apiv1)

		// 注册评论路由
		RegisterCommentsRouter(apiv1)
	}

	r.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	return r
}

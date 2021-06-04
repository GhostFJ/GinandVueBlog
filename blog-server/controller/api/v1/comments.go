package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 评论模块

// 获取单个
func GetComment(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistCommentByID(id) {
			data = models.GetComment(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// 获取多个
func GetComments(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var commentId int = -1
	if arg := c.Query("comment_id"); arg != "" {
		commentId = com.StrTo(arg).MustInt()
		maps["comment_id"] = commentId

		valid.Min(commentId, 1, "comment_id").Message("评论ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetComments(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetCommentTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增
func AddComments(c *gin.Context) {
	articleId := com.StrTo(c.Query("article_id")).MustInt()
	userId := com.StrTo(c.Query("user_id")).MustInt()

	content := c.Query("content")
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Min(articleId, 1, "article_id").Message("文章ID必须大于0")
	valid.Min(userId, 1, "user_id").Message("用户ID必须大于0")

	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(articleId) {
			data := make(map[string]interface{})
			data["article_id"] = articleId
			data["user_id"] = userId

			data["content"] = content
			data["created_by"] = createdBy

			models.AddComment(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改
func EditComments(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()

	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	valid.Min(id, 1, "id").Message("ID必须大于0")

	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        if models.ExistCommentByID(id) {
			data := models.Comment{
				Content: content,
				ModifiedBy: modifiedBy,
			}
			// data := make(map[string]interface {})

			// if content != "" {
			// 	data["content"] = content
			// }

			// data["modified_by"] = modifiedBy

			models.EditComment(id, data)
			code = e.SUCCESS
            
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })

}

// 删除
func DeleteComments(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("ID必须大于0")

    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        if models.ExistCommentByID(id) {
            models.DeleteComment(id)
            code = e.SUCCESS
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

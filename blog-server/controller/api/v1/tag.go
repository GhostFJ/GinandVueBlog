package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"net/http"
)

// @Summary 文章标签列表
// @Tags 标签管理
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} gin.Context.JSON
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增文章标签
// @Tags 标签管理
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query string true "CreatedBy"
// @Security ApiKeyAuth
// @Success 200 {string} gin.Context.JSON
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var tag models.Tag

	// ---> 绑定数据
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	name := tag.Name
	state := tag.State
	createdBy := tag.CreatedBy

	// name := c.Query("name")
	// state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	// createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// @Summary 编辑文章标签
// @Tags 标签管理
// @Produce  json
// @Param id path int true "id"
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Security ApiKeyAuth
// @Success 200 {string} gin.Context.JSON
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	//id := c.Param("id")
	id := com.StrTo(c.Param("id")).MustInt()

	var tag models.Tag

	// ---> 绑定数据
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	name := tag.Name
	state := tag.State
	modifiedBy := tag.ModifiedBy

	// name := c.Query("name")
	// modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	// var state int = -1
	// if arg := c.Query("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	// }
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// @Summary 删除文章标签
// @Tags 标签管理
// @Produce  json
// @Param id path int true "id"
// @Security ApiKeyAuth
// @Success 200 {string} gin.Context.JSON
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

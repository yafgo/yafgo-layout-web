package backend

import (
	"yafgo/yafgo-layout/internal/handler"

	"github.com/gin-gonic/gin"
)

// DmsDataHandler dms数据管理
type DmsDataHandler interface {
	TableList(ctx *gin.Context)
	TableCreate(ctx *gin.Context)
	TableDetail(ctx *gin.Context)
	TableUpdate(ctx *gin.Context)
	TableDelete(ctx *gin.Context)

	DataList(ctx *gin.Context)
	DataCreate(ctx *gin.Context)
	DataDetail(ctx *gin.Context)
	DataUpdate(ctx *gin.Context)
	DataDelete(ctx *gin.Context)
}

func NewDmsDataHandler(
	handler *handler.Handler,
) DmsDataHandler {
	return &dmsDataHandler{
		Handler: handler,
	}
}

type dmsDataHandler struct {
	*handler.Handler
}

// List implements DmsDataHandler.
//
//	@Summary	Dms数据库 list
//	@Description
//	@Tags		DMS数据管理
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases [get]
func (h *dmsDataHandler) TableList(ctx *gin.Context) {
	list := []gin.H{
		{"id": 0, "name": "default"},
	}
	total := len(list)

	h.Resp().Success(ctx, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get implements DmsDataHandler.
//
//	@Summary	Dms数据库 查询单条
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [get]
func (h *dmsDataHandler) TableDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Create implements DmsDataHandler.
//
//	@Summary	Dms数据库 新增
//	@Description
//	@Tags		DMS数据管理
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases [post]
func (h *dmsDataHandler) TableCreate(ctx *gin.Context) {
	h.Resp().Success(ctx, gin.H{
		"data": "/api/",
	})
}

// Update implements DmsDataHandler.
//
//	@Summary	Dms数据库 更新
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [post]
func (h *dmsDataHandler) TableUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Delete implements DmsDataHandler.
//
//	@Summary	Dms数据库 删除
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [delete]
func (h *dmsDataHandler) TableDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// List implements DmsDataHandler.
//
//	@Summary	Dms数据库 list
//	@Description
//	@Tags		DMS数据管理
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases [get]
func (h *dmsDataHandler) DataList(ctx *gin.Context) {
	list := []gin.H{
		{"id": 0, "name": "default"},
	}
	total := len(list)

	h.Resp().Success(ctx, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get implements DmsDataHandler.
//
//	@Summary	Dms数据库 查询单条
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [get]
func (h *dmsDataHandler) DataDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Create implements DmsDataHandler.
//
//	@Summary	Dms数据库 新增
//	@Description
//	@Tags		DMS数据管理
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases [post]
func (h *dmsDataHandler) DataCreate(ctx *gin.Context) {
	h.Resp().Success(ctx, gin.H{
		"data": "/api/",
	})
}

// Update implements DmsDataHandler.
//
//	@Summary	Dms数据库 更新
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [post]
func (h *dmsDataHandler) DataUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Delete implements DmsDataHandler.
//
//	@Summary	Dms数据库 删除
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [delete]
func (h *dmsDataHandler) DataDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

package backend

import (
	"yafgo/yafgo-layout/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type DmsDbHandler interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetTables(ctx *gin.Context)
	GetTable(ctx *gin.Context)
}

func NewDmsDbHandler(
	handler *handler.Handler,
) DmsDbHandler {
	return &dmsDbHandler{
		Handler: handler,
	}
}

type dmsDbHandler struct {
	*handler.Handler
}

// List implements DmsDbHandler.
//
//	@Summary	Dms数据库 list
//	@Description
//	@Tags		DMS数据管理
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases [get]
func (h *dmsDbHandler) List(ctx *gin.Context) {
	list := []gin.H{
		{"id": 0, "name": "default"},
	}
	total := len(list)

	h.Resp().Success(ctx, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get implements DmsDbHandler.
//
//	@Summary	Dms数据库 查询单条
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [get]
func (h *dmsDbHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Create implements DmsDbHandler.
//
//	@Summary	Dms数据库 新增
//	@Description
//	@Tags		DMS数据管理
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases [post]
func (h *dmsDbHandler) Create(ctx *gin.Context) {
	h.Resp().Success(ctx, gin.H{
		"data": "/api/",
	})
}

// Update implements DmsDbHandler.
//
//	@Summary	Dms数据库 更新
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [post]
func (h *dmsDbHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Delete implements DmsDbHandler.
//
//	@Summary	Dms数据库 删除
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id} [delete]
func (h *dmsDbHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// GetTables implements DmsDbHandler.
//
//	@Summary	Dms数据库 获取数据表
//	@Description
//	@Tags		DMS数据管理
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id}/tables [get]
func (h *dmsDbHandler) GetTables(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	tables, err := h.SvcDms.GetTablesByDBID(ctx, id)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}
	h.Resp().Success(ctx, tables)
}

// GetTables implements DmsDbHandler.
//
//	@Summary	Dms数据库 获取数据表
//	@Description
//	@Tags		DMS数据管理
//	@Param		id			path		int		true	"id"
//	@Param		tableName	path		string	true	"表名"
//	@Success	200			{object}	any		"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/dms/databases/{id}/tables/{tableName} [get]
func (h *dmsDbHandler) GetTable(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	tableName := ctx.Param("tableName")

	tables, err := h.SvcDms.GetTableColumnsByDBID(ctx, id, tableName)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}
	h.Resp().Success(ctx, tables)
}

package backend

import (
	"yafgo/yafgo-layout/internal/handler"
	"yafgo/yafgo-layout/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type MenuHandler interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Menus(ctx *gin.Context)
}

func NewMenuHandler(
	handler *handler.Handler,
) MenuHandler {
	return &menuHandler{
		Handler: handler,
	}
}

type menuHandler struct {
	*handler.Handler
}

// List implements MenuHandler.
//
//	@Summary	Menu list
//	@Description
//	@Tags		后台
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/menu/menus [get]
func (h *menuHandler) List(ctx *gin.Context) {
	list, err := h.SvcMenu.GetList(ctx)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}
	h.Resp().Success(ctx, list)
}

// Detail implements MenuHandler.
//
//	@Summary	Menu 查询单条
//	@Description
//	@Tags		后台
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/menu/menus/{id} [get]
func (h *menuHandler) Detail(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	menu, err := h.SvcMenu.GetByID(ctx, id)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}
	h.Resp().Success(ctx, menu)
}

// Create implements MenuHandler.
//
//	@Summary	Menu 新增
//	@Description
//	@Tags		后台
//	@Param		data	body		model.Menu	true	"请求参数"
//	@Success	200		{object}	any			"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/menu/menus [post]
func (h *menuHandler) Create(ctx *gin.Context) {
	item := new(model.Menu)
	if err := ctx.ShouldBindJSON(item); err != nil {
		h.ParamError(ctx, err)
		return
	}

	err := h.SvcMenu.CreateOne(ctx, item)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}

	h.Resp().Success(ctx, item)
}

// Update implements MenuHandler.
//
//	@Summary	Menu 更新
//	@Description
//	@Tags		后台
//	@Param		id		path		int			true	"id"
//	@Param		data	body		model.Menu	true	"请求参数"
//	@Success	200		{object}	any			"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/menu/menus/{id} [post]
func (h *menuHandler) Update(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	item := new(model.Menu)
	if err := ctx.ShouldBindJSON(item); err != nil {
		h.ParamError(ctx, err)
		return
	}
	item.ID = id
	_, err := h.SvcMenu.UpdateOne(ctx, item)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}

	h.Resp().Success(ctx, item)
}

// Delete implements MenuHandler.
//
//	@Summary	Menu 删除
//	@Description
//	@Tags		后台
//	@Param		id	path		int	true	"id"
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/menu/menus/{id} [delete]
func (h *menuHandler) Delete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))

	rows, err := h.SvcMenu.DelByID(ctx, id)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}

	h.Resp().Success(ctx, gin.H{
		"rows": rows,
	})
}

// Delete implements MenuHandler.
//
//	@Summary	后台菜单
//	@Description
//	@Tags		后台
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/menu [get]
func (h *menuHandler) Menus(ctx *gin.Context) {
	menus := []model.Route{
		{
			Path: "/dashboard",
			Name: "dashboard",
			Meta: model.RouteMeta{
				Locale:       "menu.server.dashboard",
				RequiresAuth: true,
				Icon:         "icon-dashboard",
				Order:        1,
			},
			Children: []model.Route{
				{
					Path: "workplace",
					Name: "Workplace",
					Meta: model.RouteMeta{
						Locale:       "menu.server.workplace",
						RequiresAuth: true,
						Icon:         "icon-dashboard",
						Order:        1,
					},
				},
				{
					Path: "https://arco.design",
					Name: "arcoWebsite",
					Meta: model.RouteMeta{
						Locale:       "menu.arcoWebsite",
						RequiresAuth: true,
						Icon:         "icon-dashboard",
						Order:        2,
					},
				},
				{
					Path: "https://arco.design",
					Name: "arcoWebsite1",
					Meta: model.RouteMeta{
						Locale:       "测试",
						RequiresAuth: true,
						Icon:         "icon-dashboard",
						Order:        3,
					},
				},
			},
		},
	}
	h.Resp().Success(ctx, menus)
}

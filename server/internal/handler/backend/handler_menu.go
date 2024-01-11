package backend

import (
	"yafgo/yafgo-layout/internal/handler"
	"yafgo/yafgo-layout/internal/model"

	"github.com/gin-gonic/gin"
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
//	@Summary		Menu list
//	@Description
//	@Tags			API
//	@Success		200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security		ApiToken
//	@Router			/admin/menu/menus [get]
func (h *menuHandler) List(ctx *gin.Context) {
	h.Resp().Success(ctx, gin.H{
		"data": "/api/",
	})
}

// Detail implements MenuHandler.
//
//	@Summary		Menu 查询单条
//	@Description
//	@Tags			API
//	@Success		200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security		ApiToken
//	@Router			/admin/menu/menus/{id} [get]
func (h *menuHandler) Detail(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Create implements MenuHandler.
//
//	@Summary		Menu 新增
//	@Description
//	@Tags			API
//	@Success		200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security		ApiToken
//	@Router			/admin/menu/menus [post]
func (h *menuHandler) Create(ctx *gin.Context) {
	h.Resp().Success(ctx, gin.H{
		"data": "/api/",
	})
}

// Update implements MenuHandler.
//
//	@Summary		Menu 更新
//	@Description
//	@Tags			API
//	@Success		200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security		ApiToken
//	@Router			/admin/menu/menus/{id} [post]
func (h *menuHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Delete implements MenuHandler.
//
//	@Summary		Menu 删除
//	@Description
//	@Tags			API
//	@Success		200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security		ApiToken
//	@Router			/admin/menu/menus/{id} [delete]
func (h *menuHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	h.Resp().Success(ctx, gin.H{
		"data": id,
	})
}

// Delete implements MenuHandler.
//
//	@Summary		后台菜单
//	@Description
//	@Tags			API
//	@Success		200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security		ApiToken
//	@Router			/admin/menu [get]
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

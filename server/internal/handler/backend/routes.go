package backend

import (
	"yafgo/yafgo-layout/internal/handler"
	"yafgo/yafgo-layout/internal/middleware"
	"yafgo/yafgo-layout/pkg/jwtutil"
	"yafgo/yafgo-layout/pkg/sys/ylog"

	"github.com/gin-gonic/gin"
)

type Router struct {
	logger *ylog.Logger
	Jwt    *jwtutil.JwtUtil

	// handler
	hdl           *handler.Handler
	indexHandler  IndexHandler
	menuHandler   MenuHandler
	systemHandler SystemHandler
	dmsDbHandler  DmsDbHandler
}

func NewRouter(
	logger *ylog.Logger,
	jwt *jwtutil.JwtUtil,
	hdl *handler.Handler,
) *Router {
	return &Router{
		logger: logger,
		hdl:    hdl,
		Jwt:    jwt,

		// handler
		indexHandler:  NewIndexHandler(hdl),
		menuHandler:   NewMenuHandler(hdl),
		systemHandler: NewSystemHandler(hdl),
		dmsDbHandler:  NewDmsDbHandler(hdl),
	}
}

func (p *Router) Register(router *gin.Engine) {
	mwAuth := middleware.JWTAuth(p.Jwt, false)
	// mwAuthForce := middleware.JWTAuth(p.Jwt, true)

	rApi := router.Group("/api/admin", mwAuth)

	{
		r := rApi.Group("")

		// index
		r.GET("/index", p.indexHandler.Index)

		// 菜单
		r.GET("/menu/menus", p.menuHandler.List)
		r.POST("/menu/menus", p.menuHandler.Create)
		r.GET("/menu/menus/:id", p.menuHandler.Detail)
		r.POST("/menu/menus/:id", p.menuHandler.Update)
		r.DELETE("/menu/menus/:id", p.menuHandler.Delete)
		// 我的菜单
		r.GET("/menu", p.menuHandler.Menus)

		// 系统配置
		r.GET("/system/cfg", p.systemHandler.ShowCfg)
		r.GET("/system/cfg_in_redis", p.systemHandler.GetCfgInRedis)
		r.POST("/system/cfg_in_redis", p.systemHandler.SetCfgInRedis)
	}

	{
		// dms数据管理
		r := rApi.Group("/dms")

		// 数据源管理
		r.GET("/databases", p.dmsDbHandler.List)
		r.POST("/databases", p.dmsDbHandler.Create)
		r.GET("/databases/:id", p.dmsDbHandler.Get)
		r.POST("/databases/:id", p.dmsDbHandler.Update)
		r.DELETE("/databases/:id", p.dmsDbHandler.List)
		r.GET("/databases/:id/tables", p.dmsDbHandler.GetTables)
		r.GET("/databases/:id/tables/:tableName", p.dmsDbHandler.GetTable)

		// 数据表管理
		r.GET("/tables")
		r.POST("/tables")
		r.GET("/tables/:tid")
		r.POST("/tables/:tid")
		r.DELETE("/tables/:tid")

		// 数据管理
		r.GET("/tables/:tid/data")
		r.POST("/tables/:tid/data")
		r.GET("/tables/:tid/data/:id")
		r.POST("/tables/:tid/data/:id")
		r.DELETE("/tables/:tid/data/:id")
	}
}

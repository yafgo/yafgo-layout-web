package frontend

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
	hdl          *handler.Handler
	indexHandler IndexHandler
	webHandler   WebHandler
	userHandler  UserHandler
}

func NewRouter(
	logger *ylog.Logger,
	hdl *handler.Handler,
	jwt *jwtutil.JwtUtil,
) *Router {
	return &Router{
		logger: logger,
		hdl:    hdl,
		Jwt:    jwt,

		// handler
		indexHandler: NewIndexHandler(hdl),
		webHandler:   NewWebHandler(hdl),
		userHandler:  NewUserHandler(hdl),
	}
}

func (p *Router) Register(router *gin.Engine) {
	mwAuth := middleware.JWTAuth(p.Jwt, false)
	// mwAuthForce := middleware.JWTAuth(p.Jwt, true)

	// 前台web
	{
		r := router.Group("/")

		// root
		router.GET("", p.webHandler.Root)
		// index
		r.GET("/index", p.webHandler.Index)
	}

	rApi := router.Group("/api", mwAuth)

	// 前台接口 [v1]
	{
		r := rApi.Group("/v1")

		// index
		rApi.GET("", p.indexHandler.Root)
		r.GET("", p.indexHandler.Index)
		r.GET("todo", todo)

		// user
		{
			r.POST("/user/register/username", p.userHandler.RegisterByUsername)
			r.POST("/user/login/username", p.userHandler.LoginByUsername)
			// 我的
			r.GET("/user/info", p.userHandler.GetProfile)
		}
	}

}

func todo(c *gin.Context) {
	reqUri := c.Request.RequestURI
	c.JSON(200, gin.H{"todo": reqUri, "success": true, "code": 0})
}

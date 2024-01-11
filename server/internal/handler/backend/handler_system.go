package backend

import (
	"yafgo/yafgo-layout/internal/handler"

	"github.com/gin-gonic/gin"
)

type SystemHandler interface {
	// ShowCfg 获取当前生效的配置
	ShowCfg(ctx *gin.Context)
	// 获取redis中的配置
	GetCfgInRedis(ctx *gin.Context)
	// 更新redis中的配置
	SetCfgInRedis(ctx *gin.Context)
}

func NewSystemHandler(
	handler *handler.Handler,
) SystemHandler {
	return &systemHandler{
		Handler: handler,
	}
}

type systemHandler struct {
	*handler.Handler
}

// ShowCfg implements SystemHandler.
//
//	@Summary	获取当前生效的配置
//	@Description
//	@Tags		后台
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/system/cfg [get]
func (h *systemHandler) ShowCfg(ctx *gin.Context) {
	settings := h.Cfg.AllSettings()
	h.Resp().Success(ctx, settings)
}

// GetCfgInRedis implements SystemHandler.
//
//	@Summary	获取redis中的配置
//	@Description
//	@Tags		后台
//	@Success	200	{object}	any	"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/system/cfg_in_redis [get]
func (h *systemHandler) GetCfgInRedis(ctx *gin.Context) {
	content, err := h.Cfg.GetRedisContent(ctx)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}
	h.Resp().Success(ctx, content)
}

type ReqCfgInRedis struct {
	Content string `json:"content"`
}

// SetCfgInRedis implements SystemHandler.
//
//	@Summary	更新redis中的配置
//	@Description
//	@Tags		后台
//	@Param		data	body		ReqCfgInRedis	true	"请求体"
//	@Success	200		{object}	any				"{"code": 200, "data": [...]}"
//	@Security	ApiToken
//	@Router		/admin/system/cfg_in_redis [post]
func (h *systemHandler) SetCfgInRedis(ctx *gin.Context) {
	req := new(ReqCfgInRedis)
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.ParamError(ctx, err)
		return
	}

	err := h.Cfg.SetRedisContent(ctx, req.Content)
	if err != nil {
		h.Resp().Error(ctx, err)
		return
	}
	h.Resp().Success(ctx)
}

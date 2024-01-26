package service

import (
	"yafgo/yafgo-layout/internal/database/query"
	"yafgo/yafgo-layout/pkg/jwtutil"
	"yafgo/yafgo-layout/pkg/sys/ycfg"
	"yafgo/yafgo-layout/pkg/sys/ylog"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service struct {
	Cfg    *ycfg.Config
	Logger *ylog.Logger
	Jwt    *jwtutil.JwtUtil
	DB     *gorm.DB
	Redis  *redis.Client
	Q      *query.Query
}

func NewService(
	cfg *ycfg.Config,
	logger *ylog.Logger,
	jwt *jwtutil.JwtUtil,
	db *gorm.DB,
	rdb *redis.Client,
	q *query.Query,
) *Service {
	return &Service{
		Cfg:    cfg,
		Logger: logger,
		Jwt:    jwt,
		DB:     db,
		Redis:  rdb,
		Q:      q,
	}
}

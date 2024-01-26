package service

import (
	"context"
	"yafgo/yafgo-layout/internal/database/model"
)

type DmsService interface {
	GetDatabases(ctx context.Context) error
	GetTablesByDBID(ctx context.Context, id int64) (tables []*model.SchemaTable, err error)
	GetTableColumnsByDBID(ctx context.Context, id int64, table string) (cols []*model.SchemaColumn, err error)
}

func NewDmsService(service *Service) DmsService {
	return &dmsService{
		Service: service,
	}
}

type dmsService struct {
	*Service

	// 默认数据库名称
	defaultDBName string
}

// GetTablesByDBID implements DmsService.
func (s *dmsService) GetTablesByDBID(ctx context.Context, id int64) (tables []*model.SchemaTable, err error) {
	var currDB = s.GetDefaultDBName(ctx)
	if id > 0 {
		currDB = "mysql"
	}
	tables, err = s.GetTables(ctx, currDB)
	return
}

// GetTableColumnsByDBID implements DmsService.
func (s *dmsService) GetTableColumnsByDBID(ctx context.Context, id int64, table string) (cols []*model.SchemaColumn, err error) {
	var currDB = s.GetDefaultDBName(ctx)
	if id > 0 {
		currDB = "mysql"
	}
	cols, err = s.GetTableColumns(ctx, currDB, table)
	return
}

// GetDatabases implements DmsService.
func (s *dmsService) GetDatabases(ctx context.Context) error {
	panic("unimplemented")
}

// ------------------------------

// GetTables implements DmsService.
func (s *dmsService) GetTables(ctx context.Context, dbName string) (tables []*model.SchemaTable, err error) {
	tables = []*model.SchemaTable{}
	sql := "SELECT * FROM information_schema.TABLES WHERE TABLE_SCHEMA=?"
	tx := s.DB.Raw(sql, dbName)
	err = tx.Scan(&tables).Error
	return
}

// GetTableColumns implements DmsService.
func (s *dmsService) GetTableColumns(ctx context.Context, dbName string, tableName string) (cols []*model.SchemaColumn, err error) {
	cols = []*model.SchemaColumn{}
	sql := "SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=? AND TABLE_NAME=?"
	tx := s.DB.Raw(sql, dbName, tableName)
	err = tx.Scan(&cols).Error
	return
}

func (s *dmsService) GetDefaultDBName(ctx context.Context) string {
	if s.defaultDBName == "" {
		rows := []*struct {
			DB string `gorm:"db" json:"db"`
		}{}
		sql := "SELECT database() AS `db`"
		tx := s.DB.Raw(sql)
		err := tx.Scan(&rows).Error
		if err != nil {
			return s.defaultDBName
		}
		if len(rows) > 0 {
			s.defaultDBName = rows[0].DB
		}
	}
	return s.defaultDBName
}

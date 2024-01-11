package service

import (
	"context"
	"yafgo/yafgo-layout/internal/model"
)

type MenuService interface {
	GetByID(ctx context.Context, id int64) (*model.Menu, error)
	GetList(ctx context.Context) ([]*model.Menu, error)
	CreateOne(ctx context.Context, item *model.Menu) error
	UpdateOne(ctx context.Context, item *model.Menu) (rows int64, err error)
	DelByID(ctx context.Context, id int64) (rows int64, err error)
}

func NewMenuService(service *Service) MenuService {
	return &menuService{
		Service: service,
	}
}

type menuService struct {
	*Service
}

// GetByID implements MenuService.
func (s *menuService) GetByID(ctx context.Context, id int64) (*model.Menu, error) {
	do := s.Q.Menu.WithContext(ctx)
	return do.GetByID(id)
}

// GetList implements MenuService.
func (s *menuService) GetList(ctx context.Context) ([]*model.Menu, error) {
	q := s.Q.Menu
	do := q.WithContext(ctx)
	list, err := do.Order(q.Order).Find()
	return list, err
}

// CreateOne implements MenuService.
func (s *menuService) CreateOne(ctx context.Context, item *model.Menu) error {
	q := s.Q.Menu
	do := q.WithContext(ctx)
	return do.Create(item)
}

// UpdateOne implements MenuService.
func (s *menuService) UpdateOne(ctx context.Context, item *model.Menu) (rows int64, err error) {
	q := s.Q.Menu
	do := q.WithContext(ctx)
	info, err := do.Where(q.ID.Eq(item.ID)).Updates(item)
	if err != nil {
		return
	}
	return info.RowsAffected, info.Error
}

// DelByID implements MenuService.
func (s *menuService) DelByID(ctx context.Context, id int64) (rows int64, err error) {
	q := s.Q.Menu
	do := q.WithContext(ctx)
	info, err := do.Where(q.ID.Eq(id)).Delete()
	if err != nil {
		return
	}
	return info.RowsAffected, info.Error
}

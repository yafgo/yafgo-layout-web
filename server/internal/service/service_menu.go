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
	GetRoutes(ctx context.Context) (routes []*model.Route, err error)
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

// GetRoutes implements MenuService.
func (s *menuService) GetRoutes(ctx context.Context) (routes []*model.Route, err error) {
	q := s.Q.Menu
	do := q.WithContext(ctx)
	list, err := do.Order(q.Order).Find()
	if err != nil || len(list) == 0 {
		return
	}

	// 构造前端所需路由结构
	var rMap = make(map[int64]*model.Route)
	var subRoutes = make([]*model.Route, 0)
	for _, v := range list {
		var item = &model.Route{
			ID:       v.ID,
			Pid:      v.Pid,
			Path:     v.Path,
			Name:     v.Name,
			Redirect: v.Redirect,
			Meta: model.RouteMeta{
				Icon:  v.Icon,
				Title: v.Label,
				Order: v.Order,
			},
		}
		rMap[v.ID] = item
		if v.Pid == 0 {
			routes = append(routes, item)
		} else {
			subRoutes = append(subRoutes, item)
		}
	}
	for _, v := range subRoutes {
		if pRoute := rMap[v.Pid]; pRoute != nil {
			if pRoute.Children == nil {
				pRoute.Children = make([]*model.Route, 0)
			}
			pRoute.Children = append(pRoute.Children, v)
		}
	}

	return
}

package model

import "encoding/json"

type Route struct {
	ID       int64     `json:"id"`                 // 菜单id
	Pid      int64     `json:"pid"`                // 菜单父id
	Path     string    `json:"path"`               //
	Name     string    `json:"name"`               // 路由名, 必须和前端路由定义一致
	Title    string    `json:"title"`              // 菜单名称
	Redirect string    `json:"redirect"`           //
	Status   *int32    `json:"status"`             // 状态, 1-启用,0-禁用
	Meta     RouteMeta `json:"meta,omitempty"`     //
	Children []*Route  `json:"children,omitempty"` //
}

type RouteMeta struct {
	RequiresAuth       bool     `json:"requiresAuth,omitempty"`
	Icon               string   `json:"icon"`                         // 菜单图标
	Title              string   `json:"title"`                        // 菜单名称
	HideInMenu         bool     `json:"hideInMenu,omitempty"`         // If true, it is not displayed in the side menu
	HideChildrenInMenu bool     `json:"hideChildrenInMenu,omitempty"` // if set true, the children are not displayed in the side menu
	ActiveMenu         string   `json:"activeMenu,omitempty"`         // 显示高亮的路由路径
	Breadcrumb         bool     `json:"breadcrumb,omitempty"`         // 是否在breadcrumb中显示
	NoAffix            bool     `json:"noAffix,omitempty"`            // if set true, the tag will not affix in the tab-bar
	NoCache            bool     `json:"noCache,omitempty"`            // if set true, the page will not be cached
	Roles              []string `json:"roles,omitempty"`              // Controls roles that have access to the page
	Order              int32    `json:"order"`                        // 一级菜单排序
}

func (p *Route) String() string {
	bs, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bs)
}

func (p *RouteMeta) String() string {
	bs, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bs)
}

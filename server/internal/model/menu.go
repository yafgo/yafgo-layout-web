package model

type Route struct {
	ID       int64     `json:"id"`                 // 菜单id
	Pid      int64     `json:"pid"`                // 菜单父id
	Path     string    `json:"path"`               //
	Name     string    `json:"name"`               // 路由名, 必须和前端路由定义一致
	Redirect string    `json:"redirect,omitempty"` //
	Meta     RouteMeta `json:"meta,omitempty"`     //
	Children []*Route  `json:"children,omitempty"` //
}

type RouteMeta struct {
	RequiresAuth       bool     `json:"requiresAuth,omitempty"`
	Icon               string   `json:"icon"`                         // The icon show in the side menu
	Title              string   `json:"title"`                        // The title name show in side menu and breadcrumb
	HideInMenu         bool     `json:"hideInMenu,omitempty"`         // If true, it is not displayed in the side menu
	HideChildrenInMenu bool     `json:"hideChildrenInMenu,omitempty"` // if set true, the children are not displayed in the side menu
	ActiveMenu         string   `json:"activeMenu,omitempty"`         // if set name, the menu will be highlighted according to the name you set
	Order              int32    `json:"order"`                        // Sort routing menu items. If set key, the higher the value, the more forward it is
	NoAffix            bool     `json:"noAffix,omitempty"`            // if set true, the tag will not affix in the tab-bar
	IgnoreCache        bool     `json:"ignoreCache,omitempty"`        // if set true, the page will not be cached
	Roles              []string `json:"roles,omitempty"`              // Controls roles that have access to the page
}

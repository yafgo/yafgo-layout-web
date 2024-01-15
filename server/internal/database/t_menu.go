package database

import "gorm.io/gen"

func init() {
	addApplyFuncs(func(g *gen.Generator) {
		// 菜单表
		g.ApplyInterface(
			func(querierMenu) {}, // 附加生成该接口定义的 Query 方法
			g.GenerateModelAs(
				"t_menu",                         // 数据库表名
				"Menu",                           // 生成的 model 名
				gen.FieldIgnore(fieldsIgnore...), // 生成的 model 忽略相应字段
				gen.FieldType("meta", "*string"),
				gen.FieldType("status", "*int32"),
			),
		)
	})
}

type querierMenu interface {
	// where(id=@id)
	GetByID(id int64) (*gen.T, error)
}

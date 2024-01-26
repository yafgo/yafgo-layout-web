package database

import "gorm.io/gen"

func init() {
	addApplyFuncs(func(g *gen.Generator) {
		g.ApplyBasic(g.GenerateModelAs(
			"t_dms_data", // 数据库表名
			"DmsData",    // 生成的 model 名
		))

		g.ApplyBasic(g.GenerateModelAs(
			"t_dms_data_column", // 数据库表名
			"DmsDataColumn",     // 生成的 model 名
		))
	})
}

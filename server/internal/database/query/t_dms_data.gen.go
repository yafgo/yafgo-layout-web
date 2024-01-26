// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"yafgo/yafgo-layout/internal/database/model"
)

func newDmsData(db *gorm.DB, opts ...gen.DOOption) dmsData {
	_dmsData := dmsData{}

	_dmsData.dmsDataDo.UseDB(db, opts...)
	_dmsData.dmsDataDo.UseModel(&model.DmsData{})

	tableName := _dmsData.dmsDataDo.TableName()
	_dmsData.ALL = field.NewAsterisk(tableName)
	_dmsData.ID = field.NewInt64(tableName, "id")
	_dmsData.DbID = field.NewInt64(tableName, "db_id")
	_dmsData.TableName_ = field.NewString(tableName, "table_name")
	_dmsData.Name = field.NewString(tableName, "name")
	_dmsData.Comment = field.NewString(tableName, "comment")
	_dmsData.Filter = field.NewString(tableName, "filter")
	_dmsData.Sql = field.NewString(tableName, "sql")
	_dmsData.Sort = field.NewString(tableName, "sort")
	_dmsData.CreatedAt = field.NewTime(tableName, "created_at")
	_dmsData.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dmsData.DeletedAt = field.NewField(tableName, "deleted_at")

	_dmsData.fillFieldMap()

	return _dmsData
}

type dmsData struct {
	dmsDataDo dmsDataDo

	ALL        field.Asterisk
	ID         field.Int64
	DbID       field.Int64  // 数据库id
	TableName_ field.String // 真实表名
	Name       field.String // 显示表名
	Comment    field.String // 表注释
	Filter     field.String // 自定义过滤条件
	Sql        field.String // 自定义SQL
	Sort       field.String // 自定义排序
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (d dmsData) Table(newTableName string) *dmsData {
	d.dmsDataDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dmsData) As(alias string) *dmsData {
	d.dmsDataDo.DO = *(d.dmsDataDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dmsData) updateTableName(table string) *dmsData {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.DbID = field.NewInt64(table, "db_id")
	d.TableName_ = field.NewString(table, "table_name")
	d.Name = field.NewString(table, "name")
	d.Comment = field.NewString(table, "comment")
	d.Filter = field.NewString(table, "filter")
	d.Sql = field.NewString(table, "sql")
	d.Sort = field.NewString(table, "sort")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")

	d.fillFieldMap()

	return d
}

func (d *dmsData) WithContext(ctx context.Context) IDmsDataDo { return d.dmsDataDo.WithContext(ctx) }

func (d dmsData) TableName() string { return d.dmsDataDo.TableName() }

func (d dmsData) Alias() string { return d.dmsDataDo.Alias() }

func (d *dmsData) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dmsData) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["id"] = d.ID
	d.fieldMap["db_id"] = d.DbID
	d.fieldMap["table_name"] = d.TableName_
	d.fieldMap["name"] = d.Name
	d.fieldMap["comment"] = d.Comment
	d.fieldMap["filter"] = d.Filter
	d.fieldMap["sql"] = d.Sql
	d.fieldMap["sort"] = d.Sort
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
}

func (d dmsData) clone(db *gorm.DB) dmsData {
	d.dmsDataDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dmsData) replaceDB(db *gorm.DB) dmsData {
	d.dmsDataDo.ReplaceDB(db)
	return d
}

type dmsDataDo struct{ gen.DO }

type IDmsDataDo interface {
	gen.SubQuery
	Debug() IDmsDataDo
	WithContext(ctx context.Context) IDmsDataDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDmsDataDo
	WriteDB() IDmsDataDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDmsDataDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDmsDataDo
	Not(conds ...gen.Condition) IDmsDataDo
	Or(conds ...gen.Condition) IDmsDataDo
	Select(conds ...field.Expr) IDmsDataDo
	Where(conds ...gen.Condition) IDmsDataDo
	Order(conds ...field.Expr) IDmsDataDo
	Distinct(cols ...field.Expr) IDmsDataDo
	Omit(cols ...field.Expr) IDmsDataDo
	Join(table schema.Tabler, on ...field.Expr) IDmsDataDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDmsDataDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDmsDataDo
	Group(cols ...field.Expr) IDmsDataDo
	Having(conds ...gen.Condition) IDmsDataDo
	Limit(limit int) IDmsDataDo
	Offset(offset int) IDmsDataDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDmsDataDo
	Unscoped() IDmsDataDo
	Create(values ...*model.DmsData) error
	CreateInBatches(values []*model.DmsData, batchSize int) error
	Save(values ...*model.DmsData) error
	First() (*model.DmsData, error)
	Take() (*model.DmsData, error)
	Last() (*model.DmsData, error)
	Find() ([]*model.DmsData, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DmsData, err error)
	FindInBatches(result *[]*model.DmsData, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DmsData) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDmsDataDo
	Assign(attrs ...field.AssignExpr) IDmsDataDo
	Joins(fields ...field.RelationField) IDmsDataDo
	Preload(fields ...field.RelationField) IDmsDataDo
	FirstOrInit() (*model.DmsData, error)
	FirstOrCreate() (*model.DmsData, error)
	FindByPage(offset int, limit int) (result []*model.DmsData, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDmsDataDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dmsDataDo) Debug() IDmsDataDo {
	return d.withDO(d.DO.Debug())
}

func (d dmsDataDo) WithContext(ctx context.Context) IDmsDataDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dmsDataDo) ReadDB() IDmsDataDo {
	return d.Clauses(dbresolver.Read)
}

func (d dmsDataDo) WriteDB() IDmsDataDo {
	return d.Clauses(dbresolver.Write)
}

func (d dmsDataDo) Session(config *gorm.Session) IDmsDataDo {
	return d.withDO(d.DO.Session(config))
}

func (d dmsDataDo) Clauses(conds ...clause.Expression) IDmsDataDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dmsDataDo) Returning(value interface{}, columns ...string) IDmsDataDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dmsDataDo) Not(conds ...gen.Condition) IDmsDataDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dmsDataDo) Or(conds ...gen.Condition) IDmsDataDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dmsDataDo) Select(conds ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dmsDataDo) Where(conds ...gen.Condition) IDmsDataDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dmsDataDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDmsDataDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dmsDataDo) Order(conds ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dmsDataDo) Distinct(cols ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dmsDataDo) Omit(cols ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dmsDataDo) Join(table schema.Tabler, on ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dmsDataDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dmsDataDo) RightJoin(table schema.Tabler, on ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dmsDataDo) Group(cols ...field.Expr) IDmsDataDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dmsDataDo) Having(conds ...gen.Condition) IDmsDataDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dmsDataDo) Limit(limit int) IDmsDataDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dmsDataDo) Offset(offset int) IDmsDataDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dmsDataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDmsDataDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dmsDataDo) Unscoped() IDmsDataDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dmsDataDo) Create(values ...*model.DmsData) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dmsDataDo) CreateInBatches(values []*model.DmsData, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dmsDataDo) Save(values ...*model.DmsData) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dmsDataDo) First() (*model.DmsData, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DmsData), nil
	}
}

func (d dmsDataDo) Take() (*model.DmsData, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DmsData), nil
	}
}

func (d dmsDataDo) Last() (*model.DmsData, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DmsData), nil
	}
}

func (d dmsDataDo) Find() ([]*model.DmsData, error) {
	result, err := d.DO.Find()
	return result.([]*model.DmsData), err
}

func (d dmsDataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DmsData, err error) {
	buf := make([]*model.DmsData, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dmsDataDo) FindInBatches(result *[]*model.DmsData, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dmsDataDo) Attrs(attrs ...field.AssignExpr) IDmsDataDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dmsDataDo) Assign(attrs ...field.AssignExpr) IDmsDataDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dmsDataDo) Joins(fields ...field.RelationField) IDmsDataDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dmsDataDo) Preload(fields ...field.RelationField) IDmsDataDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dmsDataDo) FirstOrInit() (*model.DmsData, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DmsData), nil
	}
}

func (d dmsDataDo) FirstOrCreate() (*model.DmsData, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DmsData), nil
	}
}

func (d dmsDataDo) FindByPage(offset int, limit int) (result []*model.DmsData, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dmsDataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dmsDataDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dmsDataDo) Delete(models ...*model.DmsData) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dmsDataDo) withDO(do gen.Dao) *dmsDataDo {
	d.DO = *do.(*gen.DO)
	return d
}
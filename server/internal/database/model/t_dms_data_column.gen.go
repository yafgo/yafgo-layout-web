// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDmsDataColumn = "t_dms_data_column"

// DmsDataColumn mapped from table <t_dms_data_column>
type DmsDataColumn struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	DataID        int64          `gorm:"column:data_id;not null;comment:数据表id" json:"data_id"`
	TableName_    string         `gorm:"column:table_name;not null;comment:真实表名" json:"table_name"`
	Name          string         `gorm:"column:name;not null;comment:显示列名" json:"name"`
	ColumnName    string         `gorm:"column:column_name;not null;comment:真实列名" json:"column_name"`
	ColumnType    string         `gorm:"column:column_type;not null;comment:列类型" json:"column_type"`
	DataType      string         `gorm:"column:data_type;not null;comment:数据类型" json:"data_type"`
	Order         int32          `gorm:"column:order;not null;comment:字段排序" json:"order"`
	IsNullable    bool           `gorm:"column:is_nullable;not null;default:1;comment:可为null, 0:否,1:是" json:"is_nullable"`
	CharMaxLength int32          `gorm:"column:char_max_length;not null;comment:字段最大字符长度" json:"char_max_length"`
	Comment       string         `gorm:"column:comment;not null;comment:字段注释" json:"comment"`
	Settings      string         `gorm:"column:settings;comment:字段配置" json:"settings"`
	CreatedAt     time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName DmsDataColumn's table name
func (*DmsDataColumn) TableName() string {
	return TableNameDmsDataColumn
}

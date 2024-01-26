package model

import (
	"time"
)

type SchemaTable struct {
	TABLE_CATALOG   string    `gorm:"TABLE_CATALOG" json:"TABLE_CATALOG"`
	TABLE_SCHEMA    string    `gorm:"TABLE_SCHEMA" json:"TABLE_SCHEMA"`
	TABLE_NAME      string    `gorm:"TABLE_NAME" json:"TABLE_NAME"`
	TABLE_TYPE      string    `gorm:"TABLE_TYPE" json:"TABLE_TYPE"`
	ENGINE          string    `gorm:"ENGINE" json:"ENGINE"`
	VERSION         int64     `gorm:"VERSION" json:"VERSION"`
	ROW_FORMAT      string    `gorm:"ROW_FORMAT" json:"ROW_FORMAT"`
	TABLE_ROWS      int64     `gorm:"TABLE_ROWS" json:"TABLE_ROWS"`
	AVG_ROW_LENGTH  int64     `gorm:"AVG_ROW_LENGTH" json:"AVG_ROW_LENGTH"`
	DATA_LENGTH     int64     `gorm:"DATA_LENGTH" json:"DATA_LENGTH"`
	MAX_DATA_LENGTH int64     `gorm:"MAX_DATA_LENGTH" json:"MAX_DATA_LENGTH"`
	INDEX_LENGTH    int64     `gorm:"INDEX_LENGTH" json:"INDEX_LENGTH"`
	DATA_FREE       int64     `gorm:"DATA_FREE" json:"DATA_FREE"`
	AUTO_INCREMENT  int64     `gorm:"AUTO_INCREMENT" json:"AUTO_INCREMENT"`
	TABLE_COLLATION string    `gorm:"TABLE_COLLATION" json:"TABLE_COLLATION"`
	CHECKSUM        int64     `gorm:"CHECKSUM" json:"CHECKSUM"`
	CREATE_OPTIONS  string    `gorm:"CREATE_OPTIONS" json:"CREATE_OPTIONS"`
	TABLE_COMMENT   string    `gorm:"TABLE_COMMENT" json:"TABLE_COMMENT"`
	CREATE_TIME     time.Time `gorm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATE_TIME     time.Time `gorm:"UPDATE_TIME" json:"UPDATE_TIME"`
	CHECK_TIME      time.Time `gorm:"CHECK_TIME" json:"CHECK_TIME"`
}

type SchemaColumn struct {
	TABLE_CATALOG            string `gorm:"TABLE_CATALOG" json:"TABLE_CATALOG"`
	TABLE_SCHEMA             string `gorm:"TABLE_SCHEMA" json:"TABLE_SCHEMA"`
	TABLE_NAME               string `gorm:"TABLE_NAME" json:"TABLE_NAME"`
	COLUMN_NAME              string `gorm:"COLUMN_NAME" json:"COLUMN_NAME"`
	ORDINAL_POSITION         int64  `gorm:"ORDINAL_POSITION" json:"ORDINAL_POSITION"`
	COLUMN_DEFAULT           string `gorm:"COLUMN_DEFAULT" json:"COLUMN_DEFAULT"`
	IS_NULLABLE              string `gorm:"IS_NULLABLE" json:"IS_NULLABLE"`
	DATA_TYPE                string `gorm:"DATA_TYPE" json:"DATA_TYPE"`
	CHARACTER_MAXIMUM_LENGTH int64  `gorm:"CHARACTER_MAXIMUM_LENGTH" json:"CHARACTER_MAXIMUM_LENGTH"`
	CHARACTER_OCTET_LENGTH   int64  `gorm:"CHARACTER_OCTET_LENGTH" json:"CHARACTER_OCTET_LENGTH"`
	NUMERIC_PRECISION        int64  `gorm:"NUMERIC_PRECISION" json:"NUMERIC_PRECISION"`
	NUMERIC_SCALE            int64  `gorm:"NUMERIC_SCALE" json:"NUMERIC_SCALE"`
	DATETIME_PRECISION       int64  `gorm:"DATETIME_PRECISION" json:"DATETIME_PRECISION"`
	CHARACTER_SET_NAME       string `gorm:"CHARACTER_SET_NAME" json:"CHARACTER_SET_NAME"`
	COLLATION_NAME           string `gorm:"COLLATION_NAME" json:"COLLATION_NAME"`
	COLUMN_TYPE              string `gorm:"COLUMN_TYPE" json:"COLUMN_TYPE"`
	COLUMN_KEY               string `gorm:"COLUMN_KEY" json:"COLUMN_KEY"`
	EXTRA                    string `gorm:"EXTRA" json:"EXTRA"`
	PRIVILEGES               string `gorm:"PRIVILEGES" json:"PRIVILEGES"`
	COLUMN_COMMENT           string `gorm:"COLUMN_COMMENT" json:"COLUMN_COMMENT"`
}

package model

import (
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/setting"
	"database/sql"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type KeyData struct {
	InstanceKey sql.NullInt64
	RemoetKey   sql.NullString
	FileName    sql.NullString
	Modality    sql.NullString
	Ip          sql.NullString
	SVirtualDir sql.NullString
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.DBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.MaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.MaxIdleConns)
	db.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}

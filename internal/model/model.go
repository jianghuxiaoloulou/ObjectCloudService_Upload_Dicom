package model

import (
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/setting"
	"database/sql"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type KeyData struct {
	rec_id, instance_key                                           sql.NullInt64
	jpgfile, dcmfile, ip, virpath, dttm_time                       sql.NullString
	jpgstatus, dcmstatus, objtect_time, Nfsdcmstatus, Nfsjpgstatus sql.NullInt16
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

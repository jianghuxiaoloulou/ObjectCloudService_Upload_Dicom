package model

import (
	"WowjoyProject/ObjectCloudService_Upload_Dicom/global"
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/general"
)

// 自动上传公有云数据
func GetUploadPublicData() {
	if global.RunStatus {
		global.Logger.Info("上次获取的数据没有消耗完，等待消耗完，再获取数据....")
		return
	}
	global.RunStatus = true
	global.Logger.Info("******自动上传公有云数据******")
	GetData()
	global.RunStatus = false
}

// 自动上传私有云数据
func GetUploadPrivateData() {
	if global.RunStatus {
		global.Logger.Info("上次获取的数据没有消耗完，等待消耗完，再获取数据....")
		return
	}
	global.Logger.Info("******自动上传私有云数据******")
	global.RunStatus = true
	GetData()
	global.RunStatus = false
}

func GetData() {
	sql := ""
	switch global.ObjectSetting.OBJECT_Store_Type {
	case global.PublicCloud:
		sql = `select fr.instance_key,fr.dcm_file_name_remote from file_remote fr 
		where 1= 1
		and fr.dcm_file_exist = 1
		and fr.dcm_file_exist_obs_cloud = 0
		and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ?
		limit ?;`
	case global.PrivateCloud:
		sql = `select fr.instance_key,fr.dcm_file_name_remote from file_remote fr 
		where 1= 1
		and fr.dcm_file_exist = 1
		and fr.dcm_file_exist_obs_local = 0
		and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ?
		limit ?;`
	}
	err := global.ReadDBEngine.Ping()
	if err != nil {
		global.Logger.Error("ReadDBEngine.ping() err: ", err)
		global.ReadDBEngine.Close()
		global.ReadDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	rows, err := global.ReadDBEngine.Query(sql, global.ObjectSetting.OBJECT_TIME, global.GeneralSetting.MaxTasks)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := KeyData{}
		err = rows.Scan(&key.InstanceKey, &key.RemoetKey)
		if err != nil {
			global.Logger.Error("rows.Scan error: ", err)
			continue
		}
		// 获取文件路径
		info := GetFileInfo(key.InstanceKey.Int64)
		if info.FileName == "" {
			// 异常数据不需要处理，更新为错误数据
			UpdateLocalStatus(key.InstanceKey.Int64)
			continue
		}
		// 判断数据是否是上传数据
		var dataFlag bool
		switch global.ObjectSetting.UploadImgFlag {
		case "001":
			if info.Modality != "US" && info.Modality != "ES" {
				dataFlag = true
			}
		case "010":
			if info.Modality == "US" {
				dataFlag = true
			}
		case "011":
			if info.Modality != "ES" {
				dataFlag = true
			}
		case "100":
			if info.Modality == "ES" {
				dataFlag = true
			}
		case "101":
			if info.Modality != "US" {
				dataFlag = true
			}
		case "110":
			if info.Modality == "US" && info.Modality == "ES" {
				dataFlag = true
			}
		case "111":
			dataFlag = true
		default:
			dataFlag = true
		}
		if !dataFlag {
			global.Logger.Info("数据上传设置为：", global.ObjectSetting.UploadImgFlag, "该数据不需要上传处理,更新文件状态为4,数据key: ", key.InstanceKey.Int64)
			UpdateLocalStatus(key.InstanceKey.Int64)
			continue
		}
		filekey, filepath := general.GetFilePath(info.FileName, info.Ip, info.SVirtualDir)

		data := global.ObjectData{
			InstanceKey: key.InstanceKey.Int64,
			FileKey:     filekey,
			FilePath:    filepath,
			Type:        global.DCM,
			Count:       1,
		}
		global.ObjectDataChan <- data
	}
}

func GetFileInfo(instancekey int64) (info global.FileInfo) {
	sql := `select ins.file_name,s.modality,sl.ip,sl.s_virtual_dir 
	from instance ins 
	left join study s on ins.study_key = s.study_key 
	left join study_location sl on sl.n_station_code = ins.location_code 
	where ins.instance_key = ?;`
	err := global.ReadDBEngine.Ping()
	if err != nil {
		global.Logger.Error("ReadDBEngine.ping() err: ", err)
		global.ReadDBEngine.Close()
		global.ReadDBEngine, _ = NewDBEngine(global.DatabaseSetting)
		return
	}
	row := global.ReadDBEngine.QueryRow(sql, instancekey)
	key := KeyData{}
	err = row.Scan(&key.FileName, &key.Modality, &key.Ip, &key.SVirtualDir)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	info = global.FileInfo{
		FileName:    key.FileName.String,
		Modality:    key.Modality.String,
		Ip:          key.Ip.String,
		SVirtualDir: key.SVirtualDir.String,
	}
	return
}

// 更新异常的DCM字段
func UpdateLocalStatus(key int64) {
	sql := ""
	switch global.ObjectSetting.OBJECT_Store_Type {
	case global.PublicCloud:
		sql = `update file_remote fr set fr.dcm_file_exist_obs_cloud = 4 where fr.instance_key = ?;`
	case global.PrivateCloud:
		sql = `update file_remote fr set fr.dcm_file_exist_obs_local = 4 where fr.instance_key = ?;`
	}
	err := global.WriteDBEngine.Ping()
	if err != nil {
		global.Logger.Error("WriteDBEngine.ping() err: ", err)
		global.WriteDBEngine.Close()
		global.WriteDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	global.WriteDBEngine.Exec(sql, key)
}

// 更新不存在的JPG字段
func UpdateLocalJPGStatus(key int64) {
	sql := ""
	switch global.ObjectSetting.OBJECT_Store_Type {
	case global.PublicCloud:
		sql = `update file_remote fr set fr.img_file_exist_obs_cloud = 4 where fr.instance_key = ?;`
	case global.PrivateCloud:
		sql = `update file_remote fr set fr.img_file_exist_obs_local = 4 where fr.instance_key = ?;`
	}
	err := global.WriteDBEngine.Ping()
	if err != nil {
		global.Logger.Error("WriteDBEngine.ping() err: ", err)
		global.WriteDBEngine.Close()
		global.WriteDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	global.WriteDBEngine.Exec(sql, key)
}

// 上传数据后更新数据库
func UpdateUplaod(key int64, filetype global.FileType, remotekey string, status bool) {
	// 获取更新时时间
	err := global.WriteDBEngine.Ping()
	if err != nil {
		global.Logger.Error("WriteDBEngine.ping() err: ", err)
		global.WriteDBEngine.Close()
		global.WriteDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	switch global.ObjectSetting.OBJECT_Store_Type {
	case global.PublicCloud:
		switch filetype {
		case global.DCM:
			if status {
				global.Logger.Info("***公有云DCM数据上传成功，更新状态*** ", key)
				sql := `update file_remote fr set fr.dcm_file_exist_obs_cloud = ?,fr.dcm_location_code_obs_cloud = ?,fr.dcm_update_time_obs_cloud = now(),fr.dcm_file_name_remote = ? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 1, global.ObjectSetting.OBJECT_Upload_Success_Code, remotekey, key)
			} else {
				global.Logger.Info("***公有云DCM数据上传失败，更新状态*** ", key)
				sql := `update file_remote fr set fr.dcm_file_exist_obs_cloud = ? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 2, key)
			}
		case global.JPG:
			if status {
				global.Logger.Info("***公有云JPG数据上传成功，更新状态*** ", key)
				sql := `update file_remote fr set fr.img_file_exist_obs_cloud = ?,fr.img_update_time_obs_cloud = now(),fr.img_file_name_remote=? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 1, remotekey, key)
			} else {
				global.Logger.Info("***公有云JPG数据上传失败，更新状态*** ", key)
				sql := `update file_remote fr set fr.img_file_exist_obs_cloud = ? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 2, key)
			}
		}
	case global.PrivateCloud:
		switch filetype {
		case global.DCM:
			if status {
				global.Logger.Info("***私有云DCM数据上传成功，更新状态*** ", key)
				sql := `update file_remote fr set fr.dcm_file_exist_obs_local = ?,fr.dcm_location_code_obs_local = ?,fr.dcm_update_time_obs_local = now(),fr.dcm_file_name_remote = ? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 1, global.ObjectSetting.OBJECT_Upload_Success_Code, remotekey, key)
			} else {
				global.Logger.Info("***私有云DCM数据上传失败，更新状态*** ", key)
				sql := `update file_remote fr set fr.dcm_file_exist_obs_local = ? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 2, key)
			}
		case global.JPG:
			if status {
				global.Logger.Info("***私有云JPG数据上传成功，更新状态*** ", key)
				sql := `update file_remote fr set fr.img_file_exist_obs_local = ?,fr.img_update_time_obs_local = now(),fr.img_file_name_remote=? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 1, remotekey, key)
			} else {
				global.Logger.Info("***私有云JPG数据上传失败，更新状态*** ", key)
				sql := `update file_remote fr set fr.img_file_exist_obs_local = ? where fr.instance_key = ?;`
				global.WriteDBEngine.Exec(sql, 2, key)
			}
		}
	}
}

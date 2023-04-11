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
		switch global.ObjectSetting.UploadImgFlag {
		case "001":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and s.modality != "US" 
			and s.modality != "ES" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			order by fr.rec_id ASC 
			limit ?;`
		case "010":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and s.modality = "US" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			order by fr.rec_id ASC 
			limit ?;`
		case "100":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and s.modality = "ES" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			order by fr.rec_id ASC 
			limit ?;`
		case "011":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and s.modality != "ES" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			order by fr.rec_id ASC 
			limit ?;`
		case "101":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and s.modality != "US" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			order by fr.rec_id ASC 
			limit ?;`
		case "110":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and s.modality in ("US","ES") 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			order by fr.rec_id ASC 
			limit ?;`
		case "111":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
			fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_cloud = 0 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			AND timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		}
	case global.PrivateCloud:
		switch global.ObjectSetting.UploadImgFlag {
		case "001":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and s.modality != "US" 
			and s.modality != "ES" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		case "010":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and s.modality = "US" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		case "100":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and s.modality = "ES" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		case "011":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and s.modality != "ES" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		case "101":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and s.modality != "US" 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		case "110":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			left join study s on ins.study_key = s.study_key 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and s.modality in ("US","ES") 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		case "111":
			sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
			fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
			FROM file_remote fr 
			LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
			LEFT JOIN image im ON im.instance_key = fr.instance_key 
			LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
			where 1 =1 
			and fr.rec_id > ? 
			and fr.dcm_file_exist = 1 
			and fr.dcm_file_exist_obs_local = 0 
			and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
			and timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
			ORDER BY fr.rec_id ASC 
			LIMIT ?;`
		}
	}
	if global.ReadDBEngine.Ping() != nil {
		global.Logger.Error("ReadDBEngine.ping() err: ", global.ReadDBEngine.Ping())
		global.RunStatus = false
		return
	}
	rows, err := global.ReadDBEngine.Query(sql, global.TargetValue, global.ObjectSetting.OBJECT_TIME, global.GeneralSetting.MaxTasks)
	global.Logger.Debug("当前查询的范围的rec_id 是：", global.TargetValue)
	if err != nil {
		global.Logger.Fatal(err)
		global.RunStatus = false
		return
	}
	defer rows.Close()
	for rows.Next() {
		global.TargetValue += 1
		key := KeyData{}
		err = rows.Scan(&key.rec_id, &key.instance_key, &key.Nfsdcmstatus, &key.Nfsjpgstatus, &key.dcmstatus,
			&key.jpgstatus, &key.dcmfile, &key.jpgfile, &key.ip, &key.virpath)
		if err != nil {
			global.Logger.Fatal("rows.Scan error: ", err)
			global.RunStatus = false
			return
		}
		// DICOM
		if key.dcmstatus.Int16 == int16(global.FileNotExist) && key.Nfsdcmstatus.Int16 == int16(global.FileExist) && key.dcmfile.String != "" {
			fike_key, file_path := general.GetFilePath(key.dcmfile.String, key.ip.String, key.virpath.String)
			global.Logger.Info("需要处理的文件名：", file_path)
			data := global.ObjectData{
				InstanceKey: key.instance_key.Int64,
				FileKey:     fike_key,
				FilePath:    file_path,
				Type:        global.DCM,
				Count:       1,
			}
			global.ObjectDataChan <- data
		}
		if key.Nfsdcmstatus.Int16 == int16(global.FileExist) && key.dcmfile.String == "" {
			//异常数据不需要处理，更新为错误数据
			UpdateLocalStatus(key.instance_key.Int64)
		}
	}
}

func GetDataJPG() {
	sql := ""
	switch global.ObjectSetting.OBJECT_Store_Type {
	case global.PublicCloud:
		sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_cloud,
		fr.img_file_exist_obs_cloud,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
		FROM file_remote fr 
		LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
		LEFT JOIN image im ON im.instance_key = fr.instance_key 
		LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
		where 1 =1 
		and fr.rec_id > ? 
		and fr.img_file_exist = 1 
		and fr.img_file_exist_obs_cloud = 0 
		and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
		AND timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
		ORDER BY fr.rec_id ASC 
		LIMIT ?;`
	case global.PrivateCloud:
		sql = `SELECT fr.rec_id,fr.instance_key,fr.dcm_file_exist,fr.img_file_exist,fr.dcm_file_exist_obs_local,
		fr.img_file_exist_obs_local,ins.file_name,im.img_file_name,sl.ip,sl.s_virtual_dir 
		FROM file_remote fr 
		LEFT JOIN instance ins on fr.instance_key = ins.instance_key 
		LEFT JOIN image im ON im.instance_key = fr.instance_key 
		LEFT JOIN study_location sl on sl.n_station_code = ins.location_code 
		where 1 =1 
		and fr.rec_id > ? 
		and fr.img_file_exist = 1 
		and fr.img_file_exist_obs_local = 0 
		and timestampdiff(minute,fr.dcm_update_time_retrieve,NOW()) >= 5 
		AND timestampdiff(YEAR,fr.dcm_update_time_retrieve,now()) <= ? 
		ORDER BY fr.rec_id ASC 
		LIMIT ?;`
	}
	if global.ReadDBEngine.Ping() != nil {
		global.Logger.Error("ReadDBEngine.ping() err: ", global.ReadDBEngine.Ping())
		global.RunStatus = false
		return
	}
	rows, err := global.ReadDBEngine.Query(sql, global.TargetValue, global.ObjectSetting.OBJECT_TIME, global.GeneralSetting.MaxTasks)
	global.Logger.Debug("当前查询的范围的rec_id 是：", global.TargetValue)
	if err != nil {
		global.Logger.Fatal(err)
		global.RunStatus = false
		return
	}
	defer rows.Close()
	for rows.Next() {
		global.TargetValue += 1
		key := KeyData{}
		err = rows.Scan(&key.rec_id, &key.instance_key, &key.Nfsdcmstatus, &key.Nfsjpgstatus, &key.dcmstatus,
			&key.jpgstatus, &key.dcmfile, &key.jpgfile, &key.ip, &key.virpath)
		if err != nil {
			global.Logger.Fatal("rows.Scan error: ", err)
			global.RunStatus = false
			return
		}
		// JPG
		if key.jpgstatus.Int16 == int16(global.FileNotExist) && key.Nfsjpgstatus.Int16 == int16(global.FileExist) && key.jpgfile.String != "" {
			fike_key, file_path := general.GetFilePath(key.jpgfile.String, key.ip.String, key.virpath.String)
			global.Logger.Info("需要处理的文件名：", file_path)
			data := global.ObjectData{
				InstanceKey: key.instance_key.Int64,
				FileKey:     fike_key,
				FilePath:    file_path,
				Type:        global.JPG,
				Count:       1,
			}
			global.ObjectDataChan <- data
		}
		if key.Nfsjpgstatus.Int16 == int16(global.FileExist) && key.jpgfile.String == "" {
			//异常数据不需要处理，更新为错误数据
			UpdateLocalJPGStatus(key.instance_key.Int64)
		}
	}
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
	if global.WriteDBEngine.Ping() != nil {
		global.Logger.Error("WriteDBEngine.ping() err: ", global.ReadDBEngine.Ping())
		return
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
	if global.WriteDBEngine.Ping() != nil {
		global.Logger.Error("WriteDBEngine.ping() err: ", global.ReadDBEngine.Ping())
		return
	}
	global.WriteDBEngine.Exec(sql, key)
}

// 上传数据后更新数据库
func UpdateUplaod(key int64, filetype global.FileType, remotekey string, status bool) {
	// 获取更新时时间
	if global.WriteDBEngine.Ping() != nil {
		global.Logger.Error("WriteDBEngine.ping() err: ", global.ReadDBEngine.Ping())
		return
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

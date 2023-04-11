package global

import (
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/logger"
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	ObjectSetting   *setting.ObjectSettingS
	Logger          *logger.Logger
)

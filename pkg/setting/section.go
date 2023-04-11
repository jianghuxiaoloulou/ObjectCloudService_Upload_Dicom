package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type GeneralSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
	LogMaxSize  int
	LogMaxAge   int
	MaxThreads  int
	MaxTasks    int
	CronSpec    string
}

type DatabaseSettingS struct {
	DBConn       string
	DBType       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
}

type ObjectSettingS struct {
	OBJECT_ResId                    string
	OBJECT_AK                       string
	OBJECT_POST_Upload              string
	UPLOAD_ROOT                     string
	OBJECT_Upload_Success_Code      int
	OBJECT_Count                    int
	OBJECT_Store_Type               int
	OBJECT_TIME                     int
	File_Fragment_Size              int
	Each_Section_Size               int
	File_Split_Temp                 string
	OBJECT_Multipart_Init_URL       string
	OBJECT_Multipart_Upload_URL     string
	OBJECT_Multipart_Completion_URL string
	OBJECT_Multipart_Abortion_URL   string
	OBJECT_Interface_Type           int
	OBJECT_Temp_GET_Upload          string
	OBJECT_START_KEY                int64
	UploadImgFlag                   string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

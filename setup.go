package main

import (
	"WowjoyProject/ObjectCloudService_Upload_Dicom/global"
	"WowjoyProject/ObjectCloudService_Upload_Dicom/internal/model"
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/logger"
	"WowjoyProject/ObjectCloudService_Upload_Dicom/pkg/setting"
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	readSetup()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("General", &global.GeneralSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Object", &global.ObjectSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.GeneralSetting.LogSavePath + "/" + global.GeneralSetting.LogFileName + global.GeneralSetting.LogFileExt,
		MaxSize:   global.GeneralSetting.LogMaxSize,
		MaxAge:    global.GeneralSetting.LogMaxAge,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupReadDBEngine() error {
	var err error
	global.ReadDBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupWriteDBEngine() error {
	var err error
	global.WriteDBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func readSetup() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupReadDBEngine()
	if err != nil {
		log.Fatalf("init.setupReadDBEngine err: %v", err)
	}
	err = setupWriteDBEngine()
	if err != nil {
		log.Fatalf("init.setupWriteDBEngine err: %v", err)
	}
}

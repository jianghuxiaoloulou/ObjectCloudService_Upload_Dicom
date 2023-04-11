package general

import (
	"WowjoyProject/ObjectCloudService_Upload_Dicom/global"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type UploadFile struct {
	// 表单名称
	Name string
	// 文件全路径
	Filepath string
}

// 基础函数
// func GetFilePath(file, ip, virpath string) (path string) {
// 	path += "\\\\"
// 	path += ip
// 	path += "\\"
// 	path += virpath
// 	path += "\\"
// 	path += file
// 	return
// }

func GetFilePath(file, ip, virpath string) (key, path string) {

	key += global.ObjectSetting.UPLOAD_ROOT
	key += "\\"
	key += file
	key = strings.Replace(key, "\\", "/", -1)

	path += "\\\\"
	path += ip
	path += "\\"
	path += virpath
	path += "\\"
	path += file
	return
}

func GetFileSize(filename string) int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

// 检查文件路径
func CheckPath(path string) {
	dir, _ := filepath.Split(path)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
}

// io.copy()来复制
// 参数说明：
// src: 源文件路径
// dest: 目标文件路径
// key :值不为空是更新instance表中的localtion_code值
func CopyFile(src, dest string) (int64, error) {
	// 判断路径文件夹是否存在，不存在，创建文件夹
	CheckPath(dest)
	file1, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	file2, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file2.Close()
	return io.Copy(file2, file1)
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 大文件分段
func FileSplit(file string, size int64) map[int]string {
	global.Logger.Debug("开始执行大文件分段", file)
	fileMap := make(map[int]string)
	fileSize := GetFileSize(file)
	// 分段数
	num := int(math.Ceil(float64(fileSize) / float64(size)))
	// 分段
	fi, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		global.Logger.Error(err)
		return fileMap
	}
	b := make([]byte, size)
	var i int = 1
	for ; i <= num; i++ {

		fi.Seek(int64(i-1)*(size), 0)

		if len(b) > int((fileSize - int64(i-1)*size)) {
			b = make([]byte, fileSize-int64(i-1)*size)
		}

		fi.Read(b)
		tempdir := global.ObjectSetting.File_Split_Temp
		tempFile := tempdir + file[strings.LastIndex(file, "\\"):] + "_" + strconv.Itoa(int(i))
		global.Logger.Debug(tempFile)
		f, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			global.Logger.Error(err)
			return fileMap
		}
		f.Write(b)
		f.Close()
		fileMap[i] = tempFile
	}
	fi.Close()
	return fileMap
}

func post(reqUrl string, reqParams map[string]string, contentType string, files []UploadFile, headers map[string]string) string {
	requestBody, realContentType := getReader(reqParams, contentType, files)
	httpRequest, _ := http.NewRequest("POST", reqUrl, requestBody)
	// 添加请求头
	httpRequest.Header.Add("Content-Type", realContentType)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
	// 发送请求
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &transport,
	}

	resp, err := client.Do(httpRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)
}

func getReader(reqParams map[string]string, contentType string, files []UploadFile) (io.Reader, string) {
	if strings.Index(contentType, "json") > -1 {
		bytesData, _ := json.Marshal(reqParams)
		return bytes.NewReader(bytesData), contentType
	} else if files != nil {
		body := &bytes.Buffer{}
		// 文件写入 body
		writer := multipart.NewWriter(body)
		for _, uploadFile := range files {
			file, err := os.Open(uploadFile.Filepath)
			if err != nil {
				panic(err)
			}
			part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.Filepath))
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(part, file)
			file.Close()
		}
		// 其他参数列表写入 body
		for k, v := range reqParams {
			if err := writer.WriteField(k, v); err != nil {
				panic(err)
			}
		}
		if err := writer.Close(); err != nil {
			panic(err)
		}
		// 上传文件需要自己专用的contentType
		return body, writer.FormDataContentType()
	} else {
		urlValues := url.Values{}
		for key, val := range reqParams {
			urlValues.Set(key, val)
		}
		reqBody := urlValues.Encode()
		return strings.NewReader(reqBody), contentType
	}
}

func PostForm(reqUrl string, reqParams map[string]string, headers map[string]string) string {
	return post(reqUrl, reqParams, "application/x-www-form-urlencoded", nil, headers)
}

func PostJson(reqUrl string, reqParams map[string]string, headers map[string]string) string {
	return post(reqUrl, reqParams, "application/json", nil, headers)
}

func PostFile(reqUrl string, reqParams map[string]string, files []UploadFile, headers map[string]string) string {
	return post(reqUrl, reqParams, "multipart/form-data", files, headers)
}

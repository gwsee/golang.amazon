package http

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/axgle/mahonia"

	"git.zx-tech.net/pengfeng/amazon.seller/model"
)

var SaveActionDir = "/runtime"
var SaveActionLog = false
var SaveActionLogAns = false

// ReadText response的body为字符的时候
func ReadText(by []byte, language string) (res []map[string]string) {
	var title []string
	arr := strings.Split(string(by), "\n")
	enc := mahonia.NewDecoder(language)
	for k, v := range arr {
		if v == "" {
			continue
		}
		if k == 0 {
			title = strings.Split(v, "\t")
			for k1, v1 := range title {
				title[k1] = strings.TrimSpace(v1)
			}
			continue
		}
		item := strings.Split(v, "\t")
		m := map[string]string{}
		for k1, v1 := range title {
			st := strings.TrimSpace(item[k1])
			if enc == nil {
				m[v1] = st
			} else {
				m[enc.ConvertString(v1)] = enc.ConvertString(st)
			}
		}
		res = append(res, m)
	}
	return
}

// TransMapToStruct map 转结构体
func TransMapToStruct(mp []map[string]string, res interface{}) (err error) {
	arr, err := json.Marshal(mp)
	if err != nil {
		return
	}
	err = json.Unmarshal(arr, res)
	return
}

// TransXmlToStruct xml 转结构体
func TransXmlToStruct(by []byte, res interface{}) (err error) {
	err = xml.Unmarshal(by, res)
	return
}

func GetTime(str ...string) (day string) {
	if len(str) >= 1 {
		local, _ := time.LoadLocation(str[0])
		time.Local = local
	}
	format := "2006-01-02 15:04:05"
	if len(str) >= 2 {
		format = str[1]
	}
	day = time.Now().Format(format)
	return
}
func handErr(by []byte) (err error) {
	if strings.Contains(string(by), "?xml") && strings.Contains(string(by), "ErrorResponse") {
		var obj model.ErrorResponse
		err = TransXmlToStruct(by, &obj)
		if err != nil {
			return
		}
		err = errors.New(fmt.Sprintf("%v", obj.Error))
	}
	return
}

func saveFile(urls, method string, postData map[string]string, by []byte, ret string, errs error) (fileName string, err error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		if errs == nil {
			logFilePath = dir + SaveActionDir + "/amazon_log/" + now.Format("200601") + "/" + now.Format("02") + "/"
		} else {
			logFilePath = dir + SaveActionDir + "/amazon_log/" + now.Format("200601") + "/" + now.Format("02") + "_errs/"
		}
	}
	if _, err2 := os.Stat(logFilePath); os.IsNotExist(err2) {
		os.MkdirAll(logFilePath, 0777)
		os.Chmod(logFilePath, 0777)
	}
	{
		// 请求记录
		var saveObj []byte
		saveObj = append(saveObj, []byte(urls)...)
		saveObj = append(saveObj, '\n')
		saveObj = append(saveObj, []byte(method)...)
		saveObj = append(saveObj, '\n')
		buf, _ := json.Marshal(postData)
		saveObj = append(saveObj, buf...)
		saveObj = append(saveObj, '\n')
		if errs != nil {
			saveObj = append(saveObj, []byte(ret)...)
			saveObj = append(saveObj, '\n')
			saveObj = append(saveObj, []byte(errs.Error())...)
			saveObj = append(saveObj, '\n')
		}
		saveObj = append(saveObj, []byte("amazon_request_finished")...)
		saveObj = append(saveObj, '\n')

		logFileName := now.Format("1504") + ".log"
		fileName = logFilePath + logFileName
		fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			err = err1
			return
		}
		defer fileObj.Close()
		fileObj.Write(saveObj)
	}
	//保存文件结束！！
	if ret != "" && errs == nil && SaveActionLogAns {
		no := now.Format("1504") + fmt.Sprintf("_%v", time.Now().UnixNano())
		{
			logFileName := no + "." + ret
			fileName = logFilePath + logFileName
			fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err1 != nil {
				err = err1
				return
			}
			defer fileObj.Close()
			fileObj.Write(by)
		}
		//{
		//	logFileName := no + ".json"
		//	fileName = logFilePath + logFileName
		//	fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//	if err1 != nil {
		//		err = err1
		//		return
		//	}
		//	defer fileObj.Close()
		//}

	}
	return
}

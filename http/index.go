package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Action(method, urls, action string, postData map[string]string, headers map[string]string) (by []byte, err error) {
	var req *http.Request
	if strings.Contains(action, "json") {
		buf, _ := json.Marshal(postData)
		req, err = http.NewRequest(method, urls, bytes.NewBuffer(buf))
	} else {
		val := url.Values{}
		for k, v := range postData {
			val.Add(k, v)
		}
		if strings.ToLower(method) == "post" {
			req, err = http.NewRequest(method, urls, strings.NewReader(val.Encode()))
		} else {
			req, err = http.NewRequest(method, urls+"?"+val.Encode(), strings.NewReader(val.Encode()))
		}
	}
	if err != nil {
		err = fmt.Errorf("http.NewRequest is fail: %v", err.Error())
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do is fail: %v", err.Error())
		return
	}
	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll is fail: %v", err.Error())
		return
	}
	err = handErr(by)
	if SaveActionLog || err != nil {
		_, err1 := saveFile(urls, method, postData, by, err)
		if err1 != nil {
			fmt.Println(err1.Error(), "err1.Error")
		}
	}
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("post amazon HTTP CODE:" + fmt.Sprint(resp.StatusCode))
		return
	}
	return
}

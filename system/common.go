package system

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/widuu/gojson"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type Response struct {
	code int
	Data *gojson.Js
}

//获取结构体中字段的名称
func GetStructFieldNames(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		Logger().WithFields(logrus.Fields{
			"name": "GetStructFieldNames",
		}).Info("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	Logger().WithFields(logrus.Fields{
		"name": "GetStructFieldNames",
	}).Info("result %v\n", result)
	return result
}

//获取结构体中字段的内容
func GetStructFields(structName interface{}) []interface{} {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		Logger().WithFields(logrus.Fields{
			"name": "yancy",
		}).Info("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	structValue := reflect.ValueOf(structName)
	var result []interface{}
	for i := 0; i < fieldNum; i++ {
		rst := structValue.FieldByName(t.Field(i).Name)
		result = append(result, rst)
	}
	Logger().WithFields(logrus.Fields{
		"name": "GetStructFields",
	}).Info("result %v\n", result)
	return result
}

// Post 发起post请求
func Post(url string, content string) Response {
	var jsonStr = []byte(content)
	resp, err := http.NewRequest("POST", url,
		bytes.NewBuffer(jsonStr))
	resp.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := client.Do(resp)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)
	code, _ := strconv.Atoi(req.Status)
	res := Response{code, parasJson(string(body))}
	return res
}

//解析json
func parasJson(content string) *gojson.Js {
	return gojson.Json(content)
}

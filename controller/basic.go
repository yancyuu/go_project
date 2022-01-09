package controller

import (
	"encoding/json"
	"fmt"
	"gin_test/controller/user"
	"gin_test/system"
	"github.com/sirupsen/logrus"
	"reflect"
)

//将输入的json数组参数转换成参数

type RequestParam struct {
	timeStamp     interface{}
	version       interface{}
	interfaceName interface{}
	data          interface{}
}

//一个简单的路由映射，负责将传来的字符传转换成函数运行
type ControllerMapsType map[string]reflect.Value

var ControllerMaps ControllerMapsType

func MapRouters() ControllerMapsType {
	var ruTest user.User //注册组件

	crMap := make(ControllerMapsType, 0)
	//创建反射变量，注意这里需要传入ruTest变量的地址；
	//不传入地址就只能反射Routers静态定义的方法
	vf := reflect.ValueOf(&ruTest)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	fmt.Println("NumMethod:", mNum)
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		fmt.Println("index:", i, " MethodName:", mName)
		crMap[mName] = vf.Method(i) //<<<
	}

	return crMap
}

func Create(req string) RequestParam {
	req_temp := RequestParam{}
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(req), &dat); err == nil {
		fmt.Println("==============输入的json str 转map=======================")
		fmt.Println(dat["timeStamp"])
		fmt.Println(dat["version"])
		fmt.Println(dat["interfaceName"])
		fmt.Println(dat["data"])
	}
	system.Logger().WithFields(logrus.Fields{
		"name": "yancy",
	}).Info("post请求的输入参数%v", dat)
	req_temp.timeStamp = dat["timeStamp"]
	req_temp.version = dat["version"]
	req_temp.interfaceName = dat["interfaceName"]
	req_temp.data = dat["data"]
	return req_temp
}

func (param *RequestParam) Get(inter string) interface{} { //根据名称获取内容的函数
	var content interface{}
	switch inter {
	case "timeStamp":
		content = param.timeStamp
		break
	case "version":
		content = param.version
		break
	case "interfaceName":
		content = param.interfaceName
		break
	case "data":
		content = param.data
		break

	}
	return content
}

/**
请求包的格式：
curl "http://127.0.0.1:8080/ewew/wqwq" -d
"
{\"interfaceName\":\"UserLogin\"}
"
-H"Content-Type:application/json"


*/

func (param *RequestParam) BasicController() []reflect.Value {
	fun := (param.interfaceName).(string)
	controller_map := MapRouters() //建立映射去调用函数
	//创建带调用方法时需要传入的参数列表
	args := []reflect.Value{reflect.ValueOf((param.data).(string))}
	fmt.Printf("func %v res==%v\n", fun, args)
	system.Logger().WithFields(logrus.Fields{
		"name": "yancy",
	}).Info("传入的函数参数为 %v", args)
	//使用方法名字符串调用指定方法
	rst := controller_map[fun].Call(args)
	return rst
}

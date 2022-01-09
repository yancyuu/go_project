package main

import (
	"../ginRest/controller/user"
	"encoding/json"
	"fmt"
	"github.com/widuu/gojson"
)

func main() {
	a := user.User{}
	a.Username = "yancy"
	b, _ := json.Marshal(a)
	fmt.Printf("%v", b)
	for key, value := range gojson.Json(string(b)).Getdata() {
		fmt.Printf("key: %v , value: %v  \n", key, value)
	}
	a.AutoWechatRun()
}

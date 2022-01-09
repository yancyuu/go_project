/*
 * @Descripttion:
 * @version:
 * @Author: yancyyu
 * @Date: 2021-03-22 23:17:45
 * @LastEditors: yancyyu
 * @LastEditTime: 2022-01-09 16:18:33
 */
package main

import (
	"fmt"
	user "gin_test/controller/user"

	openwechat "github.com/eatmoreapple/openwechat"
)

func main() {
	a := user.User{}
	a.Username = "yancy"
	b := user.User{}
	b.Username = "mm"
	var bots [2]*openwechat.Bot
	bots[0] = a.RegistNewBot()
	bots[1] = b.RegistNewBot()
	fmt.Println("dsadad")
	user.AutoWechatLoging(bots)
	user.AutoWechatsBlock(bots)

}

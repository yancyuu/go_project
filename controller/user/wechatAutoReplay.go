package user

import (
	"encoding/json"
	"fmt"
	"gin_test/system"
	"runtime"

	"github.com/eatmoreapple/openwechat"
	"github.com/widuu/gojson"
)

var apiKey = []string{"51a1337c2b9d4031b89460045ddec3b7", "c6118b8873b34a4c86f437dff774dd80",
	"0326ee3b436540ceabd1884207700eb5",
	"6d95987adc144da0a19883d15afe905c", "00535128abb0472ca48ca7e801cf7b0b"}

var apiUrl = "https://www.tuling123.com/openapi/api"

// Param /**声明请求api的请求值
type Param struct {
	key    string
	info   string
	userid string
}

// Value /**声明返回的文本等值
type Value struct {
	code int
	text string
	url  string
}

func (p *Param) AutoPackResponse(content system.Response) Value {
	value := Value{}
	apiCode := content.Data.Get("code").ToInt()
	value.code = apiCode
	switch apiCode {
	case 100000:
		value.text = content.Data.Get("text").Tostring()
	case 40007:
		value.text = "你请求的内容为空"
	default:
		value.text = ""
	}
	return value
}

func (p *Param) PopRightValue(info string) Value {
	param := make(map[string]interface{})
	content := system.Post(apiUrl, info)
	value := p.AutoPackResponse(content)
	index := 0
	if value.code != 100000 {
		for {
			index += 1
			for k, v := range gojson.Json(info).Getdata() {
				if k == "key" && index < len(apiKey) {
					param["key"] = apiKey[index]
				} else if index < len(apiKey) {
					param[k] = v
				}
			}
			b, _ := json.Marshal(param)
			content = system.Post(apiUrl, string(b))
			value := p.AutoPackResponse(content)
			if value.code == 100000 {
				return value
			}
			if index >= len(apiKey) {
				return value
			}
		}
	}
	return value
}

func (u *User) AutoPackMsg() func(*openwechat.Message) {

	param := new(Param)
	param.userid = u.Username
	param.key = apiKey[0]
	Handler := func(msg *openwechat.Message) {
		fmt.Println("msg.Content" + msg.Content)
		param.info = msg.Content
		b, _ := json.Marshal(param)
		fmt.Println(string(b))
		value := param.PopRightValue(string(b))
		text, _ := msg.ReplyText(value.text)
		fmt.Println(text)
	}
	return Handler
}

func (u *User) RegistNewBot() *openwechat.Bot {
	bot := openwechat.DefaultBot()
	// 注册消息处理函数
	bot.MessageHandler = u.AutoPackMsg()
	// 设置默认的登录回调
	// 可以设置通过该uuid获取到登录的二维码
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	return bot
}

//多用户登录
func AutoWechatLoging(bots [2]*openwechat.Bot) {
	for _, bot := range bots {
		// 登录
		fmt.Println("AutoWechatLoging")
		if err := bot.Login(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("AutoWechatLoging2")

	}
}

//多用户阻塞
func AutoWechatsBlock(bots [2]*openwechat.Bot) {
	// 阻塞主程序,直到用户退出或发生异常
	for _, bot := range bots {
		err := bot.Block()
		if err != nil {
			fmt.Println(err)
		}
		//多用户开启多线程
		runtime.Gosched()
	}
}

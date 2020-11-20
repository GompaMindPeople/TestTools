//Package webderver  这个是一个webderver  的操作包
package webderver

import (
	"TestTools/library/response"
	"fmt"
	"strings"

	"github.com/gogf/gf/net/ghttp"
	"github.com/tebeka/selenium"
)

// WebDerver 该结构体主要用于管理webderver
type WebDerver struct {
	ss *selenium.Service
}

const seleniumPath = "./resource/chromedriver.exe"

// Step 执行监控
func (wd *WebDerver) Step(r *ghttp.Request) {
	// ss := wd.ss
	step := r.GetFormString("step")
	strings.Split(step, "\n")
	fmt.Println(step)
	// wd.FindElements()

}

// New  用于创建一个 WebDerver
func (wd *WebDerver) New(r *ghttp.Request) {
	port := r.GetInt("port")
	webdever, err := selenium.NewChromeDriverService(seleniumPath, port)
	if err != nil {
		response.JsonExit(r, 1, "创建ChromDriver的时候发生错误:", err.Error())
		return
	}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd1, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		response.JsonExit(r, 1, "创建ChromDriver的远程连接时候发生错误:", err.Error())
		return
	}
	err = wd1.MaximizeWindow("")
	if err != nil {
		response.JsonExit(r, 1, "窗口最大化时,发生错误:", err.Error())
		return
	}
	wd.ss = webdever
	response.JsonExit(r, 0, "创建成功!", nil)
}

//Close  用于 关闭 webderver
func (wd *WebDerver) Close(r *ghttp.Request) {

	if wd.ss != nil {
		err := wd.ss.Stop()
		if err != nil {
			response.JsonExit(r, 1, "关闭ChromDriver的时候发生错误:", err.Error())
			return
		}
		wd.ss = nil
		response.JsonExit(r, 0, "关闭ChromDriver成功", nil)
		return
	}
	response.JsonExit(r, 0, "无ChromDriver可关闭", nil)

}

// FindElements 查找一组元素
func (wd *WebDerver) FindElements(r *ghttp.Request) {

}

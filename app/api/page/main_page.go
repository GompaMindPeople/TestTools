//Package page 这个是一个主要的包
package page

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//MainPage 主要页面
type MainPage struct{}

//Index index
func (mp *MainPage) Index(r *ghttp.Request) {
	name := r.GetString("name")
	fmt.Print("name->" + name)
	r.Response.WriteTpl("index.html", g.Map{
		"id":   123,
		"name": "john",
	})
}

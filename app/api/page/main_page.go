package page

import (
	"TestTools/library/response"

	"github.com/gogf/gf/net/ghttp"
)

type MainPage struct{}

func (mp *MainPage) Index(r *ghttp.Request) {

	response.JsonExit(r, 0, "hello")
}

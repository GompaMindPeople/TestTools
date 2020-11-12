package router

import (
	"TestTools/app/api/page"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		// ctlChat := new(chat.C)

		ctlUser := new(page.MainPage)
		// group.Middleware(middleware.CORS)
		// group.ALL("/chat", ctlChat)
		group.ALL("/index", ctlUser)
		// group.ALL("/curd/:table", new(curd.C))
		// group.Group("/", func(group *ghttp.RouterGroup) {
		// 	group.Middleware(middleware.Auth)
		// 	group.ALL("/user/profile", ctlUser, "Profile")
		// })
	})
}

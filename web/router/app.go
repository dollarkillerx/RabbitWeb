/**
 * @Author: DollarKiller
 * @Description: 前端服务
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:51 2019-10-21
 */
package router

import (
	"RabbitWeb/web/controllers"
	"github.com/dollarkillerx/erguotou"
)

func RegisterApp(app *erguotou.Engine) {
	app.LoadHTMLGlob("web/view/html/**/*")
	app.Delims("{%", "%}")

	app.Get("/", controllers.Home)
}

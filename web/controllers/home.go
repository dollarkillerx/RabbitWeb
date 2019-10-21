/**
 * @Author: DollarKiller
 * @Description: home主页
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:03 2019-10-21
 */
package controllers

import "github.com/dollarkillerx/erguotou"

func Home(ctx *erguotou.Context) {
	ctx.HTML(200, "home/home.html")
}

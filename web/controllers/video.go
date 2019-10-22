/**
 * @Author: DollarKiller
 * @Description: video 视频播放
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:28 2019-10-22
 */
package controllers

import (
	"github.com/dollarkillerx/erguotou"
	"log"
)

func Video(ctx *erguotou.Context) {
	// 获取播放列表
	name, b := ctx.PathValueString("name")
	if !b {
		// 如果获取参数错误 说明非法入侵
		log.Println("用户传入参数错误")
		ctx.Redirect(302, "/")
		return
	}

	// 数据库查询是否sql中是否存在  如果存在写入html中
	name = name
	ctx.HTML(200, "home/video.html")
}

/**
 * @Author: DollarKiller
 * @Description: 视频表
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:06 2019-10-21
 */
package datamodels

import "github.com/jinzhu/gorm"

type RabVideoList struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);comment '剧集名称'"` // 剧集名称
	VideoId int    `gorm:"type:int;index:video_id;comment 'video 表关联'"`
	M3u8    string `gorm:"type:varchar(366);comment 'm3u8 url '"`
	Img     string `gorm:"type:varchar(366);comment 'img url'"`
}

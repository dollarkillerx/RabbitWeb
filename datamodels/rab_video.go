/**
 * @Author: DollarKiller
 * @Description: 视频list 表
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:06 2019-10-21
 */
package datamodels

import "github.com/jinzhu/gorm"

type RabVideo struct {
	gorm.Model
	Name         string         `gorm:"type:varchar(100);unique_index;comment '番名'"` // 番名
	Img          string         `gorm:"type:varchar(355);comment 'img url'"`
	Introduction string         `gorm:"type:text"` // 描述
	RabVideoList []RabVideoList `gorm:"foreignkey:VideoId"`
}

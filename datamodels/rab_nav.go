/**
 * @Author: DollarKiller
 * @Description: nav 导航
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:35 2019-10-21
 */
package datamodels

import "github.com/jinzhu/gorm"

type RabNav struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);comment '导航名称'"`
	Url    string `gorm:"type:varchar(255);comment 'url地址'"`
	Weight int    `gorm:"comment '权重'"`
}

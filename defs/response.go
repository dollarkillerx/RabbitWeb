/**
 * @Author: DollarKiller
 * @Description: response 统一返回定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:32 2019-10-22
 */
package defs

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

var (
	Resp400 = Resp{Code: 400, Msg: "Error"}
)

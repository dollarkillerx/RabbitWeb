/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:02 2019-10-21
 */
package main

import (
	"RabbitWeb/config"
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	err := app.Run(
		erguotou.SetHost(config.MyConfig.App.Host),
		erguotou.SetDebug(config.MyConfig.App.Debug),
		erguotou.SetUploadSize(64<<20),
	)
	if err != nil {
		panic(err)
	}
}

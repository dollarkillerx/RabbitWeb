/**
 * @Author: DollarKiller
 * @Description: 配置文件
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:05 2019-10-21
 */
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type base struct {
	App struct {
		Host       string `yaml:"host"`
		BaseUrl    string `yaml:"baseurl"`
		Debug      bool   `yaml:"debug"`
		BaseStatic string `yaml:"-"`
		Language   string `yaml:"language"`
		S3url      string `yaml:""`
	}
	MySql struct {
		Dsn string `yaml:"dsn"`
	}
	Pgsql struct {
		Dsn     string        `yaml:"dsn"`
		MaxIdle int           `yaml:"max_idle"`
		MaxOpen int           `yaml:"max_open"`
		TimeOut time.Duration `yaml:"time_out"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	MyConfig *base
)

func init() {
	MyConfig = &base{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	MyConfig.App.BaseStatic = MyConfig.App.BaseUrl + "/static"

	if MyConfig.App.Debug {
		log.Println(MyConfig)
	}
}

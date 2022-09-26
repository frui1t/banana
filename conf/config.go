package conf

import (
	"banana/model"
	"banana/server"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	ServiceConf `ini:"kafka"`
	MysqlConf   `ini:"mysql"`
	Jwt         `ini:"jwt"`
}

type ServiceConf struct {
	AppMode      string        `ini:"AppMode"`
	HttpPort     string        `ini:"HttpPort"`
	ReadTimeout  time.Duration `ini:"ReadTimeout"`
	WriteTimeout time.Duration `ini:"WriteTimeout"`
}
type MysqlConf struct {
	DB           string `ini:"DB"`
	DbHost       string `ini:"DbHost"`
	DbPort       string `ini:"DbPort"`
	DbUser       string `ini:"DbUser"`
	DbPassword   string `ini:"DbPassword"`
	DbName       string `ini:"DbName"`
	MaxIdleConns int    `ini:"MaxIdleConns"`
	MaxOpenConns int    `ini:"MaxOpenConns"`
}

type Jwt struct {
	JwtSecret string `ini:"JwtSecret"`
}

func Init() {

	var cfg = new(Config)
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		logrus.Errorf("加载配置文件失败 err: ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MysqlConf.DbUser, cfg.MysqlConf.DbPassword, cfg.MysqlConf.DbHost, cfg.MysqlConf.DbPort, cfg.MysqlConf.DbName)
	model.Database(dsn, cfg.MysqlConf.MaxIdleConns, cfg.MaxOpenConns)

	r := server.NewRouter()
	//fmt.Println(cfg.Jwt.JwtSecret)
	r.Run(":8000")

	/*
		s := &http.Server{
			Addr:           ":8000",
			Handler:        r,
			ReadTimeout:    10,
			WriteTimeout:   10,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
	*/

}

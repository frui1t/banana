package conf

import (
	"banana/model"
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	ServiceConf `ini:"kafka"`
	MysqlConf   `ini:"mysql"`
}

type ServiceConf struct {
	AppMode      string `ini:"AppMode"`
	HttpPort     string `ini:"HttpPort"`
	ReadTimeout  string `ini:"ReadTimeout"`
	WriteTimeout string `ini:"WriteTimeout"`
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

func Init() {
	var cfg = new(Config)
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		logrus.Errorf("加载配置文件失败 err: ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MysqlConf.DbUser, cfg.MysqlConf.DbPassword, cfg.MysqlConf.DbHost, cfg.MysqlConf.DbPort, cfg.MysqlConf.DbName)
	model.Database(dsn, cfg.MysqlConf.MaxIdleConns, cfg.MaxOpenConns)
}
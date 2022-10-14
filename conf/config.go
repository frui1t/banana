package conf

import (
	"banana/model"
	"banana/server"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {

	workDir, err := os.Getwd()
	workDir += "/conf/"
	if err != nil {
		logrus.Errorf("获取配置文件目录 err: ", err)
	}
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(workDir)
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("加载配置文件失败 err: ", err)
	}
	dbHost := viper.GetString("mysql.dbHost")
	dbPort := viper.GetString("mysql.dbPort")
	dbUser := viper.GetString("mysql.dbUser")
	dbPassword := viper.GetString("mysql.dbPassword")
	dbName := viper.GetString("mysql.dbName")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	model.MysqlDatabase(dsn, 10, 10)
	model.RedisDatabase()

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

package model

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Database(dsn string, maxId, maxOpen int) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		DisableDatetimePrecision:  true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		logrus.Errorf("连接数据库失败 err: ", err)
	} else {
		logrus.Info("连接数据库成功")
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("err: ", err)
	}
	sqlDB.SetMaxIdleConns(maxId)
	sqlDB.SetMaxOpenConns(maxOpen)

	DB = db

	//migration()

}

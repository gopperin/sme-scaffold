package api

import (
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	myconfig "sme-scaffold/internal/config"
	"sme-scaffold/internal/domain/product"
)

// InitDatabase InitDatabase
func InitDatabase() (*gorm.DB, error) {

	config := mysql.Config{
		DSN:                       myconfig.Case.Database.URL, // DSN data source name
		DefaultStringSize:         255,                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                      // 根据版本自动配置
	}

	logMode := logger.Silent
	switch strings.ToLower(myconfig.Case.Database.LogMode) {
	case "info":
		logMode = logger.Info
		break
	case "warn":
		logMode = logger.Warn
		break
	case "error":
		logMode = logger.Error
		break
	// silent is default
	default:
		logMode = logger.Silent
		break
	}

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logMode,     // Log level
			Colorful:      true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(config), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db, nil
}

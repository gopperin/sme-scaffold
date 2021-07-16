package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const cmdRoot = "core"

// Setup 载入配置文件
func Setup(path string) {

	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	Application = &ApplicationStruct{
		Port: viper.GetString("application.port"),
	}

	Database.Dialect = viper.GetString("database.dialect")
	Database.Database = viper.GetString("database.database")
	Database.User = viper.GetString("database.user")
	Database.Password = viper.GetString("database.password")
	Database.Host = viper.GetString("database.host")
	Database.Port = viper.GetInt("database.port")
	Database.Charset = viper.GetString("database.charset")
	Database.MaxIdleConns = viper.GetInt("database.maxIdleConns")
	Database.MaxOpenConns = viper.GetInt("database.maxOpenConns")
	Database.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		Database.User, Database.Password, Database.Host, Database.Port, Database.Database, Database.Charset)

	CORS.Enable = viper.GetBool("cors.enable")
	CORS.AllowOrigins = viper.GetStringSlice("cors.AllowOrigins")
	CORS.AllowMethods = viper.GetStringSlice("cors.AllowMethods")
	CORS.AllowHeaders = viper.GetStringSlice("cors.AllowHeaders")
	CORS.AllowCredentials = viper.GetBool("cors.AllowCredentials")
	CORS.MaxAge = viper.GetInt("cors.MaxAge")

}

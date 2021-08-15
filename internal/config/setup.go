package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const cmdRoot = "core"

// Setup 载入配置文件
func Setup(path ...string) {

	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath(path[0])

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&Case); err != nil {
			fmt.Println(err)
		}
	})
	if err := viper.Unmarshal(&Case); err != nil {
		fmt.Println(err)
	}

	Case.Database.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		Case.Database.User, Case.Database.Password, Case.Database.Host, Case.Database.Port, Case.Database.Database, Case.Database.Charset)

}

package initialization

import (
	"douyin/global"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 初始化viper.Viper结构体以提供项目的config 命令行 > 环境变量 > 默认值
func InitializeViper(path ...string) {
	var config string
	//
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "please choose a config file")
		flag.Parse()
		//若无命令行输入
		if config == "" {
			//若无环境变量输入
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = global.ConfigDefaultFile
					fmt.Printf("using gin %s mode, config path is %s\n", gin.EnvGinMode, global.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = global.ConfigReleaseFile
					fmt.Printf("using gin %s mode, config path is %s\n", gin.EnvGinMode, global.ConfigReleaseFile)
				case gin.TestMode:
					config = global.ConfigTestFile
					fmt.Printf("using gin %s mode, config path is %s\n", gin.EnvGinMode, global.ConfigTestFile)
				}

			} else {
				//若存在环境变量输入
				config = global.ConfigEnv
				fmt.Printf("using environment variable, config path is %s\n", config)
			}
		} else {
			//若存在命令行输入
			fmt.Printf("using command line, config path is %s\n", config)
		}
	} else {
		//若存在可变参数
		config = path[0]
		fmt.Printf("using value passed by path, config path is %s\n", config)
	}

	//用viper结构读取配置
	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		//可能尽量不使用panic?思考一下怎么搞
		panic(fmt.Errorf("Fatal error in config file: %s \n", err))
	}

	if err = vip.Unmarshal(&global.SERVER_CONFIG); err != nil {
		fmt.Println(err)
	}

	//添加对于配置文件的监视，如果发生变化修改viper结构，为了以后热切换用
	vip.WatchConfig()

	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = vip.Unmarshal(&global.SERVER_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	fmt.Println("====1-viper====: viper init config success")

	global.SERVER_VIPER = vip
}

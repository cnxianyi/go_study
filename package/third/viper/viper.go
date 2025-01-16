package viper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port int  `mapstructure:"port" json:"port" yaml:"port"`
	Ipv6 bool `mapstructure:"ipv6" json:"ipv6" yaml:"ipv6"`
}

// 先在当前文件夹下 新建 config.yaml 文件,并添加
// port: 0

func ViperTest() {
	viper.SetConfigName("config")                // 配置文件名，不带扩展名
	viper.AddConfigPath("./package/third/viper") // 配置文件所在路径
	viper.SetConfigType("yaml")                  // 配置文件格式类型（yaml）

	// 设置 默认值
	viper.SetDefault("ipv6", true)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 获取配置中的值
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Ipv6: %t\n", config.Ipv6)

	go func() {
		// 使用 viper 监控文件变化
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)

			// 重新加载配置
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatalf("Error reading config file after change, %s", err)
			}

			// 获取并打印新的配置
			var updatedConfig Config
			err = viper.Unmarshal(&updatedConfig)
			if err != nil {
				log.Fatalf("Unable to decode into struct after change, %v", err)
			}

			// 打印更新后的配置值
			fmt.Printf("Updated Port: %d\n", updatedConfig.Port)
			fmt.Printf("Updated Ipv6: %t\n", updatedConfig.Ipv6)
		})

		// 开始监视配置文件变化
		viper.WatchConfig()

		// 等待文件更新
		select {} // 阻塞等待
	}()

	go func() {
		time.Sleep(1 * time.Second)
		// 模拟更新文件
		// 文件 读写模式 权限
		file, err := os.OpenFile("./package/third/viper/config.yaml", os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("读取失败")
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("读取失败")
		}

		// 模拟修改端口 +1
		// Replace(读取内容文本 , 匹配的文本 , 替换的文本 , 模式) -1: 全部替换
		newData := strings.Replace(string(data), fmt.Sprintf("%d", config.Port), fmt.Sprintf("%d", config.Port+1), -1)
		file.Seek(0, 0)             // 将文件指针移到开头写入. 即覆盖
		file.Write([]byte(newData)) // 将修改后的文件写回文件
	}()

}

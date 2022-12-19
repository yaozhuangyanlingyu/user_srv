package helper

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"github.com/spf13/cast"
)

/**
 * 读取配置中心数据
 */
func GetConsulConfig(host string, port string, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 读取配置中的的地址
		consul.WithAddress(host+":"+cast.ToString(port)),

		// 设置前缀，不设置默认前缀/micro/config
		consul.WithPrefix(prefix),

		// 是否移除前缀，这里是设置为true，表示可以不带前缀，直接获取对应配置
		consul.StripPrefix(true),
	)

	// 配置初始化
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}

	// 加载配置
	err = config.Load(consulSource)
	return config, err
}

// MysqlConfig 创建结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// ...string 用做参数
// ... 参数语法形成了可变参数的参数。它将接受零个或多个参数，并将它们作为切⽚引⽤
func GetMySQLConsulConfig(config config.Config, configStruct interface{}, key ...string) error {
	//获取配置
	err := config.Get(key...).Scan(configStruct)
	if err != nil {
		return err
	}
	return nil
}

package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"site/common/tools"
	"strings"
)

// 项目根目录名称与配置文件名称
const (
	AppConfRootDir = "conf"                        // 配置文件目录 - 默认的配置文件根目录
	AppRootDir     = "site-collaboration-platform" // 项目根目录名称
)

// 配置文件类型与文件名称
const (
	// 服务配置项
	appConfDir  = "yaml"     // 配置文件目录
	appConfFile = "app.yaml" // 配置文件名称

	// 参数配置项
	paramDir      = "yaml" // 配置文件目录
	paramConfFile = "param.yaml"

	// 密钥配置项
	keyConfDir     = "key"             // 配置文件目录
	privateKeyFile = "private_key.pem" // 私钥文件
	publicKeyFile  = "public_key.pem"  // 公钥文件
)

type AppConfig struct {
	Services map[string]*ServiceConf `yaml:"services"`
}

type ServiceConf struct {
	Name  string     `yaml:"name"`
	Host  string     `yaml:"host"`
	Port  string     `yaml:"port"`
	Mysql *MysqlConf `yaml:"mysql"`
}

type MysqlConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

// LoadAppConf 加载服务地址配置
func LoadAppConf() (*AppConfig, error) {
	var c AppConfig
	v := viper.New()
	file := filePath(appConfFile, appConfDir)
	v.SetConfigFile(file)
	v.SetConfigType(appConfDir)
	// 读取配置信息
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	// 反序列化到结构中
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

// InitServicesMap 初始化 ServicesMap
func (c *AppConfig) InitServicesMap() {
	if c.Services == nil {
		c.Services = make(map[string]*ServiceConf)
	}
}

// GetServiceConf 获取某一服务配置信息
func (c *AppConfig) GetServiceConf(s string) *ServiceConf {
	if c == nil {
		return nil
	}
	c.InitServicesMap()
	if addr, exist := c.Services[s]; exist {
		return addr
	}
	return nil
}

// filePath 动态获取配置路径 用于兼容单元测试路径
func filePath(fileName string, dirType string) string {
	pathPrefix := "." // 前缀
	path := tools.GetCurrentPath()
	dirArr := strings.Split(path, "/")
	if len(dirArr) < 1 {
		panic("filePath cannot be empty")
	}
	// 获取项目根目录下具有的目录层数
	num := 0
	for i, dirName := range dirArr {
		if reflect.DeepEqual(dirName, AppRootDir) {
			num = len(dirArr) - 1 - i
			break
		}
	}
	for i := 0; i < num; i++ {
		pathPrefix = fmt.Sprintf("%s/..", pathPrefix)
	}
	path = strings.Join([]string{pathPrefix, fmt.Sprintf("%v/%v", AppConfRootDir, dirType), fileName}, "/")
	return path
}

// GetAddress 获取服务配置地址
func (s *ServiceConf) GetAddress() string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

// GetMysqlDsn 获取 mysql 配置信息
func (m *MysqlConf) GetMysqlDsn() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.DbName)
}

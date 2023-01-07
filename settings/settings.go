package settings

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/arg"
	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

type Settings struct {
	Username   string
	Password   string
	TmdbApiKey string
}

var ctx = context.Background()
var execPath = filepath.Dir(os.Args[0]) // 可执行文件所在文件夹
var confPath string                     // 默认值=execPath/conf

var Instance Settings // 读取到的配置内容

func init() {
	flag.StringVar(&confPath, "conf-path", filepath.Join(execPath, "conf"), "config path")
	flag.Parse()
	if abs, err := filepath.Abs(confPath); err == nil {
		confPath = abs // 命令行参数传入的路径
	}
}

func Init() error {
	logs.Info(ctx, "可执行文件所在目录=%v 使用的配置文件目录=%v", execPath, confPath)
	fileName := filepath.Join(confPath, "app.yaml")
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrapf(err, "读取配置文件失败(%v)", fileName)
	}
	if err := viper.Unmarshal(&Instance); err != nil {
		return errors.Wrapf(err, "反序列化配置文件失败(%v)", fileName)
	}
	logs.Info(ctx, "配置文件读取成功: %v", arg.JSON(Instance))
	return nil
}

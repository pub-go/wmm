package settings

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/arg"
	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
	"github.com/youthlin/t"
	"github.com/youthlin/t/locale"
	"golang.org/x/text/language/display"
)

type Settings struct {
	Username   string
	Password   string
	TmdbApiKey string
}

var (
	ctx      = context.Background()
	execPath = filepath.Dir(os.Args[0]) // 可执行文件所在文件夹
	Instance Settings                   // 读取到的配置内容
)

// 命令行参数
var (
	lang     string // 默认值="" 即系统默认语言
	confPath string // 默认值=execPath/conf
)

func init() {
	flag.StringVar(&lang, "lang", "", "language")
	flag.StringVar(&confPath, "conf-path", filepath.Join(execPath, "conf"), "config path")
	flag.Parse()
	if abs, err := filepath.Abs(confPath); err == nil {
		confPath = abs // 命令行参数传入的路径
	}
}

func Init() error {
	loadI18n()
	return loadConfig()
}

func loadI18n() {
	t.Load(confPath)
	t.SetLocale(lang)                // 如果指定了语言
	t.SetLocale(t.MostMatchLocale()) // 但指定的语言不一定有 重新指定最接近的
	var supportLangs []string
	for _, locale := range t.Locales() {
		supportLangs = append(supportLangs, getLangName(locale))
	}
	logs.Info(ctx, t.T("system locale=%v. want use language=%v. used language=%v. supported language=[%v]",
		getLangName(locale.GetDefault()), lang, getLangName(t.UsedLocale()), strings.Join(supportLangs, ", "),
	))
}

func getLangName(lang string) string {
	currentTag := t.Tag(t.UsedLocale())
	tag := t.Tag(lang)
	return fmt.Sprintf("%s: %s [%v]", tag.String(), display.Self.Name(tag), display.Tags(currentTag).Name(tag))
}

func loadConfig() error {
	logs.Info(ctx, t.T("executeable file path=%v, config file path=%v", execPath, confPath))
	fileName := filepath.Join(confPath, "app.yaml")
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, t.T("Read config file failed(%v)", fileName))
	}
	if err := viper.Unmarshal(&Instance); err != nil {
		return errors.Wrap(err, t.T("Unmarshal config file failed(%v)", fileName))
	}
	logs.Info(ctx, t.T("Read config file ok: %v", arg.JSON(Instance)))
	return nil
}

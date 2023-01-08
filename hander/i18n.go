package hander

import (
	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/kv"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
	"golang.org/x/text/language"
)

const (
	HeaderAcceptLanguage = "Accept-Language"
	KeyTranslations      = "translations"
)

func I18n(c *gin.Context) {
	SetTranslations(c)
	c.Next()
}

func SetTranslations(c *gin.Context) {
	accept := c.Request.Header.Get(HeaderAcceptLanguage)
	tags, _, _ := language.ParseAcceptLanguage(accept)
	if len(tags) == 0 {
		tags = append(tags, t.Tag(t.Locale()))
	}
	var supported = t.Locales()
	_, index, _ := t.MatchTag(t.Tags(supported), tags)
	lang := supported[index]
	ctx := kv.Add(c.Request.Context(), "lang", lang)
	c.Request = c.Request.WithContext(ctx)
	logs.Info(c, "kv.Get: %v", kv.Get(c))
	logs.Info(c, t.T("client header accept-language is %v, parsed as %v. used language %v for this request",
		accept, tags, lang,
	))
	c.Set(KeyTranslations, t.L(lang))
}

func GetTranslations(c *gin.Context) *t.Translations {
	if value, ok := c.Get(KeyTranslations); ok {
		return value.(*t.Translations)
	}
	return t.Global()
}

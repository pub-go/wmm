package hander

import (
	"code.gopub.tech/wmm/model"
	"code.gopub.tech/wmm/settings"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func Login(c *gin.Context) (any, error) {
	var req LoginReq
	var t = GetTranslations(c)
	if err := c.BindJSON(&req); err != nil {
		return nil, model.ErrInvalidParams(t).WithCause(err)
	}
	if req.Username == settings.Instance.Username && req.Password == settings.Instance.Password {
		return nil, nil
	}
	return nil, model.ErrInvalidAccount(t)
}

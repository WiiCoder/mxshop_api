package global

import (
	"mxshop_api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Translator ut.Translator

	UserSrvClient proto.UserClient
)

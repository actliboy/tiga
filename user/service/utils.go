package service

import (
	"github.com/liov/tiga/protobuf/utils/errorcode"
	"github.com/liov/tiga/user/conf"
	"github.com/liov/tiga/utils/verification"
)

func LuosimaoVerify(vCode string) error {
	if err := verification.LuosimaoVerify(conf.Conf.Customize.LuosimaoVerifyURL,
		conf.Conf.Customize.LuosimaoAPIKey, vCode); err != nil {
		return errorcode.InvalidArgument.Message(err.Error())
	}
	return nil
}

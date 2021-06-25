package service

import (
	"github.com/liov/tiga/protobuf/utils/request"
	contexti "github.com/liov/tiga/tiga/context"
	"github.com/liov/tiga/tiga/pick"
	"github.com/liov/tiga/tiga/pick/_example/middle"
	"net/http"

	model "github.com/liov/tiga/protobuf/user"
	"github.com/liov/tiga/protobuf/utils/response"
)

type UserService struct{}

func (*UserService) Service() (string, string, []http.HandlerFunc) {
	return "用户相关", "/api/user", []http.HandlerFunc{middle.Log}
}

func (*UserService) Add(ctx *contexti.Ctx, req *model.SignupReq) (*response.TinyRep, error) {
	//对于一个性能强迫症来说，我宁愿它不优雅一些也不能接受每次都调用
	pick.Api(func() {
		pick.Post("").
			Title("用户注册").
			Version(2).
			CreateLog("1.0.0", "jyb", "2019/12/16", "创建").
			ChangeLog("1.0.1", "jyb", "2019/12/16", "修改测试").End()
	})

	return &response.TinyRep{Message: "测试"}, nil
}

func (*UserService) Edit(ctx *contexti.Ctx, req *model.EditReq) (*model.EditReq_EditDetails, error) {
	pick.Api(func() {
		pick.Put("/:id").
			Title("用户编辑").
			CreateLog("1.0.0", "jyb", "2019/12/16", "创建").
			Deprecated("1.0.0", "jyb", "2019/12/16", "删除").End()
	})

	return nil, nil
}

func (*UserService) Get(ctx *contexti.Ctx, req *request.Object) (*response.TinyRep, error) {
	pick.Api(func() {
		pick.Get("/:id").
			Title("用户注册").
			CreateLog("1.0.0", "jyb", "2019/12/16", "创建").End()
	})

	return &response.TinyRep{Code: uint32(req.Id), Message: "测试"}, nil
}

type StaticService struct{}

func (*StaticService) Service() (string, string, []http.HandlerFunc) {
	return "静态资源", "/api/static", nil
}

func (*StaticService) Get2(ctx *contexti.Ctx, req *model.SignupReq) (*response.TinyRep, error) {
	pick.Api(func() {
		pick.Get("/*mail").
			Title("用户注册").
			CreateLog("1.0.0", "jyb", "2019/12/16", "创建").End()
	})

	return &response.TinyRep{Message: req.Mail}, nil
}

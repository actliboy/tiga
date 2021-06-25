package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liov/tiga/protobuf/user"
	"github.com/liov/tiga/tiga"
	"github.com/liov/tiga/tiga/initialize"
	"github.com/liov/tiga/tiga/pick"
	uconf "github.com/liov/tiga/user/conf"
	udao "github.com/liov/tiga/user/dao"
	userservice "github.com/liov/tiga/user/service"
	"github.com/liov/tiga/utils/log"
	"net/http"
	"time"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
)

func main() {
	//配置初始化应该在第一位
	defer initialize.Start(uconf.Conf, udao.Dao)()

	view.RegisterExporter(&exporter.PrintExporter{})
	view.SetReportingPeriod(time.Second)
	// Register the view to collect gRPC client stats.
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Fatal(err)
	}
	pick.RegisterService(userservice.GetUserService())
	(&tiga.Server{
		//为了可以自定义中间件
		GRPCOptions: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(),
			grpc.ChainStreamInterceptor(),
			//grpc.StatsHandler(&ocgrpc.ServerHandler{})
		},
		GRPCHandle: func(gs *grpc.Server) {
			user.RegisterUserServiceServer(gs, userservice.GetUserService())
			user.RegisterOauthServiceServer(gs, userservice.GetOauthService())
		},
		GinHandle: func(app *gin.Engine) {
			_ = user.RegisterUserServiceHandlerServer(app, userservice.GetUserService())
			_ = user.RegisterOauthServiceHandlerServer(app, userservice.GetOauthService())

			app.Static("/static", "F:/upload")
			app.StaticFS("/oauth/login", http.Dir("./static/login.html"))
			pick.Gin(app, true, initialize.InitConfig.Module)
		},
	}).Start()
}

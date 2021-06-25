package grpci

import (
	"github.com/liov/tiga/utils/log"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(log.Default)
}

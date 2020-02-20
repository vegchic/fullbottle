package main

import (
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/common/log"
	"github.com/vegchic/fullbottle/config"
	"github.com/vegchic/fullbottle/user/handler"
	"github.com/micro/go-micro/v2"

	user "github.com/vegchic/fullbottle/user/proto/user"
)

func main() {
	service := micro.NewService(
		micro.Name(config.UserSrvName),
		micro.Version("latest"),
		micro.WrapHandler(common.ServiceErrorRecovery),
		micro.WrapHandler(common.ServiceLogWrapper),
	)

	service.Init()

	common.SetClient(service.Client())

	if err := user.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler)); err != nil {
		log.Fatalf(err, "RegisterUserServiceHandler failed")
	}

	if err := service.Run(); err != nil {
		log.Fatalf(err, "Service running failed")
	}
}
package cmd

import (
	"ewallet-notification/cmd/proto/notification"
	"ewallet-notification/helpers"
	"ewallet-notification/internal/api"
	"ewallet-notification/internal/repository"
	"ewallet-notification/internal/services"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	d := dependencyInject()
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7003"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()
	notification.RegisterNotificationServiceServer(s, d.EmailAPI)

	// list method

	logrus.Info("start listening grpc on port: ", helpers.GetEnv("GRPC_PORT", "7003"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}
}

type Dependency struct {
	EmailAPI *api.EmailAPI
}

func dependencyInject() Dependency {
	emailRepo := &repository.EmailRepository{
		DB: helpers.DB,
	}
	emailSvc := &services.EmailService{
		EmailRepository: emailRepo,
	}
	emailAPI := &api.EmailAPI{
		EmailService: emailSvc,
	}

	return Dependency{
		EmailAPI: emailAPI,
	}
}

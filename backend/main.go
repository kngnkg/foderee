package main

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"github.com/kngnkg/tunetrail/backend/config"
	"github.com/kngnkg/tunetrail/backend/infra/server"
	"github.com/kngnkg/tunetrail/backend/logger"

	helloworld "github.com/kngnkg/tunetrail/backend/gen/helloworld"
	user "github.com/kngnkg/tunetrail/backend/gen/user"
	"google.golang.org/grpc"
)

func main() {
	l := logger.New(&logger.LoggerOptions{
		Level: slog.LevelDebug,
	})

	cfg, err := config.New()
	if err != nil {
		l.Fatal("failed to load config.", err)
	}

	// TODO: usecaseなどのインスタンスを生成する

	s := grpc.NewServer()

	helloworldServer := server.NewHelloworldServer()
	helloworld.RegisterGreeterServer(s, helloworldServer)

	userServer := server.NewUserServer()
	user.RegisterUserServiceServer(s, userServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		l.Fatal("failed to listen.", err)
	}

	l.Info("gRPC server started.", "port", strconv.Itoa(cfg.Port))
	if err := s.Serve(lis); err != nil {
		l.Fatal("failed to serve.", err)
	}
}

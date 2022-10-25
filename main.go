package main

import (
	"context"
	"sync"

	"github.com/buker/micro-writer/handler"
	pb "github.com/buker/micro-writer/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "micro-writer"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.BeforeStart(func() error {
			logger.Infof("Starting service %s", service)
			return nil
		}),
		micro.BeforeStop(func() error {
			logger.Infof("Shutting down service %s", service)
			cancel()
			return nil
		}),
		micro.AfterStop(func() error {
			wg.Wait()
			return nil
		}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Server().Init(
		server.Wait(&wg),
	)

	ctx = server.NewContext(ctx, srv.Server())

	// Register handler
	if err := pb.RegisterMicroWriterHandler(srv.Server(), new(handler.MicroWriter)); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

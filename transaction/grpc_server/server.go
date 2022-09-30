package grpc_server

import (
	"github.com/nurtilekassankhan/meetup/transaction/pkg/config"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/handler"
	pb "github.com/nurtilekassankhan/meetup/transaction/pkg/proto/transaction"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Run() error
	Stop()
	GracefulStop()
}

type server struct {
	pb.UnimplementedUserManagementServer
	handlr handler.Handler
	serv   *grpc.Server
	lis    net.Listener
}

func NewServer(conf *config.Server, handlr handler.Handler) (Server, error) {

	serv := grpc.NewServer()

	lis, err := net.Listen(conf.Network, conf.Port)
	if err != nil {
		return nil, err
	}

	srv := &server{
		handlr: handlr,
		serv:   serv,
		lis:    lis,
	}

	pb.RegisterUserManagementServer(serv, srv)

	return srv, nil
}

func (s *server) Run() error {
	return s.serv.Serve(s.lis)
}

func (s *server) Stop() {
	s.serv.Stop()
}

func (s *server) GracefulStop() {
	s.serv.GracefulStop()
}

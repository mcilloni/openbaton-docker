package server

import (
	"net"

	"google.golang.org/grpc"

	pop "github.com/mcilloni/openbaton-docker/pop/proto"
	"github.com/openbaton/go-openbaton/util"
	log "github.com/sirupsen/logrus"
)

// Server represents the PoP service.
type Server struct {
	*log.Logger
	Config   Config
	Listener net.Listener

	svc *service
}

// New initialises a new Server from viper.
func New() (*Server, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	l := log.New()
	l.Level = cfg.LogLevel

	return &Server{
		Logger: l,
		Config: cfg,
	}, nil
}

// Close shuts down the Server.
func (s *Server) Close() error {
	tag := util.FuncName()

	s.WithFields(log.Fields{
		"tag":      tag,
		"pop-name": s.Config.PopName,
	}).Info("stopping server")

	err1 := s.svc.close()
	err2 := s.Listener.Close()

	switch {
	case err1 != nil:
		return err1

	case err2 != nil:
		return err2

	default:
		return nil
	}
}

// Serve spawns the service.
func (s *Server) Serve() error {
	tag := util.FuncName()

	s.WithFields(log.Fields{
		"tag":      tag,
		"pop-name": s.Config.PopName,
	}).Info("starting server")

	proto := s.Config.Proto
	if proto == "" {
		proto = pop.DefaultListenProtocol
	}

	laddr := s.Config.Netaddr
	if laddr == "" {
		laddr = pop.DefaultListenAddress
	}

	lis, err := net.Listen(proto, laddr)
	if err != nil {
		return err
	}

	s.Listener = lis

	s.svc, err = newService(s.Config, s.Logger)
	if err != nil {
		return err
	}

	srv := grpc.NewServer(
		grpc.StreamInterceptor(s.svc.streamInterceptor),
		grpc.UnaryInterceptor(s.svc.unaryInterceptor),
	)

	pop.RegisterPopServer(srv, s.svc)

	s.WithFields(log.Fields{
		"tag":      tag,
		"pop-name": s.Config.PopName,
	}).Info("launching gRPC server")

	if err := srv.Serve(s.Listener); err != nil {
		return err
	}

	return nil
}

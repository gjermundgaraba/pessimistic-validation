package server

import (
	"context"
	"github.com/gjermundgaraba/pessimistic-validation/core/types"
	"github.com/gjermundgaraba/pessimistic-validation/sidecar/attestors"
	"gitlab.com/tozd/go/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	types.UnimplementedSidecarServer
	listener net.Listener

	logger      *zap.Logger
	coordinator attestors.Coordinator
	grpcServer  *grpc.Server
}

var _ types.SidecarServer = &Server{}

func NewServer(logger *zap.Logger, coordinator attestors.Coordinator) *Server {
	return &Server{
		logger: logger,
		coordinator: coordinator,
	}
}

func (s *Server) Serve(listenAddr string) error {
	s.logger.Debug("server.Serve", zap.String("listenAddr", listenAddr))

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", listenAddr)
	}

	s.grpcServer = grpc.NewServer()
	types.RegisterSidecarServer(s.grpcServer, s)
	s.logger.Info("server listening", zap.String("addr", lis.Addr().String()))
	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.logger.Debug("server.Stop")

	s.grpcServer.GracefulStop()
}

func (s *Server) GetAttestation(ctx context.Context, request *types.AttestationRequest) (*types.AttestationResponse, error) {
	s.logger.Debug("server.GetLatestAttestation", zap.String("chainId", request.ChainId))

	attestation, err := s.coordinator.GetLatestAttestation(request.ChainId)
	if err != nil {
		return nil, err
	}

	return &types.AttestationResponse{
		Attestation: &attestation,
	}, nil
}

package presenceserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/slice"
	"github.com/SoroushBeigi/knowledge-game/rpcmessage/golang/presence"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
	"google.golang.org/grpc"
)

type Server struct {
	presence.UnimplementedPresenceServiceServer
	svc presenceservice.Service
}

func New(svc presenceservice.Service) *Server {
	return &Server{
		UnimplementedPresenceServiceServer: presence.UnimplementedPresenceServiceServer{},
		svc:                                svc,
	}
}

func (s Server) GetPresence(ctx context.Context, req *presence.GetPresenceRequest) (*presence.GerPresenceResponse, error) {
	resp, err := s.svc.GetPresence(ctx, dto.GetPresenceRequest{UserIDs: slice.Uint64toUint(req.GetUserIds())})
	if err != nil {
		return nil, err
	}

	return resp.ToGrpc(), nil
}

func (s Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8086))
	if err != nil {
		log.Fatal("grpc listener error")
	}

	presenceServer := Server{}
	grpcServer := grpc.NewServer()

	presence.RegisterPresenceServiceServer(grpcServer, &presenceServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("grpc serve error")
	}

}

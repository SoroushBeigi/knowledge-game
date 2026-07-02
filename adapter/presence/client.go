package presence

import (
	"context"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/protobuf"
	"github.com/SoroushBeigi/knowledge-game/pkg/slice"
	"github.com/SoroushBeigi/knowledge-game/rpcmessage/golang/presence"
	"google.golang.org/grpc"
)

type Client struct {
	client presence.PresenceServiceClient
}

func New(conn *grpc.ClientConn) *Client {

	return &Client{
		client: presence.NewPresenceServiceClient(conn),
	}
}

func (c Client) GetPresence(ctx context.Context, req dto.GetPresenceRequest) (dto.GetPresenceResponse, error) {
	resp, err := c.client.GetPresence(ctx,
		&presence.GetPresenceRequest{UserIds: slice.UintToUint64(req.UserIDs)},
	)

	if err != nil {
		return dto.GetPresenceResponse{}, err
	}

	return protobuf.ProtobufToGetPresenceResponse(resp), nil
}

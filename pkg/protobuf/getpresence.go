package protobuf

import (
	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/rpcmessage/golang/presence"
)

func GetPresenceResponseToProtobuf(g dto.GetPresenceResponse) *presence.GetPresenceResponse {
	r := &presence.GetPresenceResponse{}

	for _, item := range g.Items {
		r.Items = append(r.Items, &presence.GetPresenceItem{
			UserId:    uint64(item.UserID),
			Timestamp: item.Timestamp,
		})
	}

	return r
}

func ProtobufToGetPresenceResponse(g *presence.GetPresenceResponse) dto.GetPresenceResponse {
	r := dto.GetPresenceResponse{}

	for _, item := range g.Items {
		r.Items = append(r.Items, dto.GetPresenceItem{
			UserID:    uint(item.UserId),
			Timestamp: item.Timestamp,
		})
	}

	return r
}

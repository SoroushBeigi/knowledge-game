package dto

import "github.com/SoroushBeigi/knowledge-game/rpcmessage/golang/presence"

type GetPresenceRequest struct {
	UserIDs []uint
}

type GetPresenceResponse struct {
	Items []GetPresenceItem
}

func (gpResponse GetPresenceResponse) ToGrpc() *presence.GerPresenceResponse {
	r := &presence.GerPresenceResponse{}

	for _, item := range gpResponse.Items {
		r.Items = append(r.Items, &presence.GetPresenceItem{
			UserId:    uint64(item.UserID),
			Timestamp: item.Timestamp,
		})
	}

	return r
}

type GetPresenceItem struct {
	UserID    uint
	Timestamp int64
}

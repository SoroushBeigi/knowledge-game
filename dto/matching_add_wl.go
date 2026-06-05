package dto

import (
	"time"

	"github.com/SoroushBeigi/knowledge-game/entity"
)

type AddToWaitingListRequest struct {
	UserID   uint            `json:"user_id"`
	Category entity.Category `json:"category"`
}

type AddToWaitingListResponse struct {
	Timeout time.Duration `json:"timeout"`
}

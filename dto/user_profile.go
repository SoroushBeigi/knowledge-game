package dto

type GetProfileRequest struct {
	UserID uint `json:"id"`
}

type GetProfileResponse struct {
	Name string `json:"name"`
}
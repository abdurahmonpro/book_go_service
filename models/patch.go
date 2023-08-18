package models

type UpdatePatchRequest struct {
	Id     string                 `json:"id"`
	Status int32					  `json:"status"`
}

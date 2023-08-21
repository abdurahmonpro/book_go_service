package models

import (
	"book/genproto/book_service"
)

type UpdatePatchRequest struct {
	Id       int32                 `json:"id"`
	Updpatch book_service.BookData `json:"updatepatch"`
}

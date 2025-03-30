package input

import "mime/multipart"

type CategoryCreateInput struct {
	Name        string
	Description string
	ImageFile   *multipart.FileHeader
}

type CategoryQueryInput struct {
	Keyword *string
}

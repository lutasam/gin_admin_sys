package bo

import "mime/multipart"

type UploadImageRequest struct {
	FileHeader *multipart.FileHeader
}

type UploadImageResponse struct {
	ImageURL string `json:"image_url"`
}

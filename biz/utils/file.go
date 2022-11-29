package utils

import (
	"mime/multipart"
)

func IsCorrectImg(header *multipart.FileHeader) (bool, error) {
	//if header.Header.Get("Content-Type") != "image/png" ||
	//	header.Header.Get("Content-Type") != "image/jpeg" ||
	//	header.Header.Get("Content-Type") != "image/jpg" {
	//	return false, common.IMGFORMATERROR
	//}
	//if header.Size > common.MAXIMGSPACE {
	//	return false, common.IMGTOOLARGEERROR
	//}
	return true, nil
}

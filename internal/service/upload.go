package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/okh8609/gin_blog/pkg/upload"
)

type FileInfo struct { // response data
	Name      string
	AccessUrl string
}

func (s *Service) UploadFile(file multipart.File, fileHeader *multipart.FileHeader, fileType upload.FileType) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename) // {MD5_hash}.xxx

	if upload.CheckExt(fileName, fileType) == false {
		return nil, errors.New("file extension not supported")
	}

	if upload.CheckMaxSize(fileHeader, fileType) == false {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(dst, fileHeader); err != nil {
		return nil, err
	}

	if upload.CheckMIME(dst, fileType) == false {
		return nil, errors.New("file MIME is not supported")
	}

	if upload.CheckMaxSizeOnServer(dst, fileType) == false {
		return nil, errors.New("exceeded maximum file limit.(2)")
	}

	accessUrl := upload.GetServerUrl() + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

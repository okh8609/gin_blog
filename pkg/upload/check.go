package upload

import (
	"context"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/okh8609/gin_blog/global"
)

type FileType int

const (
	ImageFile FileType = iota + 1
	DocFile
)

func CheckExt(name string, t FileType) bool {
	ext := path.Ext(name)
	ext = strings.ToUpper(ext)

	switch t {
	case ImageFile:
		for _, allowExt := range global.App.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	case DocFile:
		for _, allowExt := range global.App.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

func CheckMaxSize(f *multipart.FileHeader, t FileType) bool {
	// Web Server也要擋
	// 上傳前就先檢查
	size := f.Size / 1024 / 1024 // MB
	global.MyLogger.Debugf(context.Background(), "# multipart.FileHeader.Size = %v Bytes.#", f.Size)

	switch t {
	case ImageFile:
		if size <= global.App.UploadImageMaxSize {
			return true
		}
	case DocFile:
		if size <= global.App.UploadDocMaxSize {
			return true
		}
	}

	return false
}

func CheckSavePath(dst string) bool {
	// 檢查檔案路徑是否存在
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

func CheckMIME(path string, t FileType) bool {
	// 檢查檔案MIME
	MIME := GetMIME(path)

	switch t {
	case ImageFile:
		for _, allowMIME := range global.App.UploadImageAllowMIME {
			if MIME == allowMIME {
				return true
			}
		}
	case DocFile:
		for _, allowMIME := range global.App.UploadDocAllowMIME {
			if MIME == allowMIME {
				return true
			}
		}
	}

	DeleteFile(path)
	return false
}

func CheckMaxSizeOnServer(path string, t FileType) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	size := f.Size() / 1024 / 1024 // get the size (MB)
	global.MyLogger.Debugf(context.Background(), "# CheckMaxSizeOnServer() f.Size = %v Bytes.#", f.Size())

	switch t {
	case ImageFile:
		if size <= global.App.UploadImageMaxSize {
			return true
		}
	case DocFile:
		if size <= global.App.UploadDocMaxSize {
			return true
		}
	}

	DeleteFile(path)
	return false
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

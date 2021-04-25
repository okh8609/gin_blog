package upload

import (
	"context"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/pkg/utils"
)

func GetMIME(path string) string {
	// open the uploaded file
	file, err := os.Open(path)
	if err != nil {
		global.MyLogger.Error(context.Background(), err)
		return "--- error ---"
	}

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)
	if err != nil {
		global.MyLogger.Error(context.Background(), err)
		return "--- error ---"
	}

	filetype := http.DetectContentType(buff)
	global.MyLogger.Infof(context.Background(), "Path: %s; Type: %s", path, filetype)

	return filetype
}

func GetFileName(name string) string {
	// timestamp-{MD5_hash}.xxx
	ext := path.Ext(name)

	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.MD5(fileName)
	fileName = fileName[0:16]

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	return timestamp + "_" + fileName + ext
}

func GetSavePath() string {
	return global.App.UploadSavePath
}

func GetServerUrl() string {
	return global.App.UploadServerUrl
}

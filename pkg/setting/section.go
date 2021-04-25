package setting

import "time"

// https://yaml2go.prasadg.dev/

// Server
type ServerSetting struct {
	RunMode                string        `yaml:"RunMode"`
	HttpPort               string        `yaml:"HttpPort"`
	ReadTimeout            time.Duration `yaml:"ReadTimeout"`
	WriteTimeout           time.Duration `yaml:"WriteTimeout"`
	HandleMethodNotAllowed bool          `yaml:"HandleMethodNotAllowed"`
}

// App
type AppSetting struct {
	DefaultPageSize int    `yaml:"DefaultPageSize"`
	MaxPageSize     int    `yaml:"MaxPageSize"`
	LogSavePath     string `yaml:"LogSavePath"`
	LogFileName     string `yaml:"LogFileName"`
	LogFileExt      string `yaml:"LogFileExt"`

	UploadSavePath       string   `yaml:"UploadSavePath"`
	UploadServerUrl      string   `yaml:"UploadServerUrl"`
	UploadImageMaxSize   int64    `yaml:"UploadImageMaxSize"`
	UploadImageAllowMIME []string `yaml:"UploadImageAllowMIME"`
	UploadImageAllowExts []string `yaml:"UploadImageAllowExts"`
	UploadDocMaxSize     int64    `yaml:"UploadDocMaxSize"`
	UploadDocAllowMIME   []string `yaml:"UploadDocAllowMIME"`
	UploadDocAllowExts   []string `yaml:"UploadDocAllowExts"`
}

// Database
type DatabaseSetting struct {
	DBType       string `yaml:"DBType"`
	UserName     string `yaml:"UserName"`
	Password     string `yaml:"Password"`
	Host         string `yaml:"Host"`
	DBName       string `yaml:"DBName"`
	TablePrefix  string `yaml:"TablePrefix"`
	Charset      string `yaml:"Charset"`
	ParseTime    bool   `yaml:"ParseTime"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
}

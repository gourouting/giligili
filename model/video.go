package model

import "github.com/jinzhu/gorm"

// Video 视频模型
type Video struct {
	gorm.Model
	Title string
	Info  string
	URL   string
}

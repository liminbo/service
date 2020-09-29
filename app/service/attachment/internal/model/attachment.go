package model

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"github.com/micro/go-micro/v2/logger"
	"github.com/pkg/errors"
	"strconv"
	"sync"
)

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

type Attachment struct {
	Aid      int32  `gorm:"column:aid;not null;AUTO_INCREMENT;PRIMARY_KEY;type:int(11)" json:"aid"`
	UserID   int32  `gorm:"column:user_id;not null;type:int(10)" json:"user_id"`
	Path     string `gorm:"column:path;not null;type:varchar(255)" json:"path"`
	Attr     string `gorm:"column:attr;not null;comment: '相关属性JSON字符串',;type:text" json:"attr"`
	ImgStore int8   `gorm:"column:img_store;not null;default:'0';comment: '0 又拍云 1 腾讯云',;type:tinyint(4)" json:"img_store"`
	Type     int8   `gorm:"column:type;not null;default:'0';comment: '类型 0 图片 1 视频',;type:tinyint(4)" json:"type"`
	Created  int64  `gorm:"column:created;not null;type:int(11)" json:"created"`
}

type ImageExt struct {
	W    int32  `json:"w"`    // 宽度
	H    int32  `json:"h"`    // 高度
	Mime string `json:"mime"` // 类型
}

type VideoExt struct {
	W     int32  `json:"w"`     // 宽度
	H     int32  `json:"h"`     // 高度
	Mime  string `json:"mime"`  // 视频类型
	Cover string `json:"cover"` // 封面图
}

// 数据库存放的 宽高都是string类型的 golang 解析报错
type FormatImageExt struct {
	W    string `json:"w"`    // 宽度
	H    string `json:"h"`    // 高度
	Mime string `json:"mime"` // 类型
}

type FormatVideoExt struct {
	W     string `json:"w"`     // 宽度
	H     string `json:"h"`     // 高度
	Mime  string `json:"mime"`  // 视频类型
	Cover string `json:"cover"` // 封面图
}

func (Attachment) TableName() string {
	return `b_attachment`
}

func (image *ImageExt) Decode(s string) (err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	_image := &FormatImageExt{}
	if err = json.Unmarshal([]byte(s), _image); err == nil {
		if w, convErr := strconv.Atoi(_image.W); convErr == nil {
			image.W = int32(w)
		}
		if h, convErr := strconv.Atoi(_image.H); convErr == nil {
			image.H = int32(h)
		}
		image.Mime = _image.Mime
		return
	}

	__image := &ImageExt{}
	if err = json.Unmarshal([]byte(s), __image); err == nil {
		image.W = __image.W
		image.H = __image.H
		image.Mime = _image.Mime
		return
	}

	logger.Warnf("ImageExt decode err:%v string:%s", err, s)
	return
}

func (image *ImageExt) Encode() (s string, err error) {
	var(
		b []byte
		json = jsoniter.ConfigCompatibleWithStandardLibrary
	)
	_image := &FormatImageExt{
		W:    strconv.Itoa(int(image.W)),
		H:    strconv.Itoa(int(image.H)),
		Mime: image.Mime,
	}

	if b, err = json.Marshal(_image); err != nil {
		logger.Warnf("ImageExt encode err:%v data:%v", err, image)
		return "", err
	}

	buf := bfPool.Get().(*bytes.Buffer)
	if _, err = buf.Write(b); err != nil{
		err = errors.Wrap(err, "ImageExt buf write")
		return
	}
	s = buf.String()
	buf.Reset()
	bfPool.Put(buf)

	return
}

func (video *VideoExt) Decode(s string) (err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	_video := &FormatVideoExt{}

	if err = json.Unmarshal([]byte(s), _video); err != nil {
		logger.Warnf("ImageExt decode err:%v string:%s", err, s)
		return
	}

	if w, convErr := strconv.Atoi(_video.W); convErr == nil {
		video.W = int32(w)
	}
	if h, convErr := strconv.Atoi(_video.H); convErr == nil {
		video.H = int32(h)
	}

	return
}

func (video *VideoExt) Encode() (s string, err error) {
	var (
		b []byte
		json = jsoniter.ConfigCompatibleWithStandardLibrary
	)
	_video := &FormatVideoExt{
		W:     strconv.Itoa(int(video.W)),
		H:     strconv.Itoa(int(video.W)),
		Mime:  video.Mime,
		Cover: video.Cover,
	}

	if b, err = json.Marshal(_video); err != nil {
		logger.Warnf("ImageExt encode err:%v data:%v", err, video)
		return "", err
	}
	s = string(b)
	return
}

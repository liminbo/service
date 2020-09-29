package dao

import (
	v1 "micro-service/app/service/attachment/api/v1"
	"micro-service/app/service/attachment/internal/model"
	"sync"
	"time"
)

var (
	atPool = sync.Pool{
		New: func() interface{} {
			return &model.Attachment{}
		},
	}
)

func (d *Dao) AttachmentByIds(ids []int32) (attachmentList []*model.Attachment, err error) {
	db := d.db.Where("aid IN (?)", ids)
	err = db.Find(&attachmentList).Error
	return
}

func (d *Dao) AddAttachment(req *v1.AddAttachmentReq) (id int32, err error){
	attachment := atPool.Get().(*model.Attachment)
	attachment.UserID = req.GetUserId()
	attachment.Path = req.GetUrl()
	attachment.Type = int8(req.GetAttachType())
	attachment.Created = time.Now().Unix()

	if imageExt := req.GetImageExt(); imageExt != nil{
		attachment.Attr = d.encodeImageExt(imageExt)
	}
	if videoExt := req.GetVideoExt(); videoExt != nil{
		attachment.Attr = d.encodeVideoExt(videoExt)
	}

	if err = d.db.Model(attachment).Create(attachment).Error; err != nil{
		return
	}
	id = attachment.Aid

	*attachment = model.Attachment{}
	atPool.Put(attachment)
	return
}

func (d *Dao) encodeImageExt(ext *v1.ImageExt) (str string){
	_imageExt := &model.ImageExt{
		W: ext.GetW(),
		H: ext.GetH(),
		Mime:ext.GetMime(),
	}
	if attr, extErr := _imageExt.Encode(); extErr == nil{
		str = attr
	}
	return
}

func (d *Dao) encodeVideoExt(ext *v1.VideoExt) (str string){
	_videoExt := &model.VideoExt{
		W:     ext.GetW(),
		H:     ext.GetH(),
		Mime:  ext.GetMime(),
		Cover: ext.GetCover(),
	}

	if attr, extErr := _videoExt.Encode(); extErr == nil{
		str = attr
	}
	return
}
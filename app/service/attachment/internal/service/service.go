package service

import (
	"context"
	"github.com/micro/go-micro/v2/logger"
	v1 "micro-service/app/service/attachment/api/v1"
	"micro-service/app/service/attachment/internal/dao"
	"micro-service/app/service/attachment/internal/model"
)

// Service service.
type Service struct {
	dao *dao.Dao
}

// New new a service and return.
func New(d *dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

// Close close the resource.
func (s *Service) Close() {
	logger.Info("close service")
}

func (s *Service) AttachmentDetailByIds(ctx context.Context, req *v1.AttachmentDetailByIdsReq, rep *v1.AttachmentDetailByIdsRep) (err error) {
	var (
		attachmentList []*model.Attachment
	)

	if attachmentList, err = s.dao.AttachmentByIds(req.GetIds()); err != nil {
		logger.Errorf("AttachmentDetailByIds err:%v", err)
		return
	}

	data := make(map[int32]*v1.Attachment)
	for _, value := range attachmentList {
		attachment := &v1.Attachment{
			Id:         value.Aid,
			Url:        value.Path,
			UserId:     value.UserID,
			AttachType: v1.AttachType(value.Type),
			CreatedAt:  int32(value.Created),
		}
		switch attachment.AttachType {
		case v1.AttachType_IMAGE:
			_imageExt := &model.ImageExt{}
			if extErr := _imageExt.Decode(value.Attr); extErr == nil {
				imageExt := &v1.ImageExt{
					W:    _imageExt.W,
					H:    _imageExt.H,
					Mime: _imageExt.Mime,
				}
				attachment.ImageExt = imageExt
			}
		case v1.AttachType_VIDEO:
			_videoExt := &model.VideoExt{}
			if extErr := _videoExt.Decode(value.Attr); extErr == nil {
				videoExt := &v1.VideoExt{
					W:     _videoExt.W,
					H:     _videoExt.H,
					Mime:  _videoExt.Mime,
					Cover: _videoExt.Cover,
				}
				attachment.VideoExt = videoExt
			}
		}
		data[value.Aid] = attachment
	}

	rep.Attachment = data
	return
}

func (s *Service) AddAttachment(ctx context.Context, req *v1.AddAttachmentReq, rep *v1.AddAttachmentReqRep) (err error) {
	var (
		id int32
	)

	if id, err = s.dao.AddAttachment(req); err != nil {
		logger.Errorf("AddAttachment err:%v", err)
		return
	}

	rep.Id = id
	return
}

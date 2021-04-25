package service

import (
	"github.com/okh8609/gin_blog/internal/model"
	"github.com/okh8609/gin_blog/pkg/app"
)

func (s *Service) CountTag(param *GetTagsParam) (int64, error) {
	return s.dao.CountTag(param.Name, param.State)
}

func (s *Service) GetTagList(param *GetTagsParam, pager *app.Pager) ([]*model.BlogTag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(param *CreateTagParam) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (s *Service) UpdateTag(param *UpdateTagParam) error {
	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (s *Service) DeleteTag(param *DeleteTagParam) error {
	return s.dao.DeleteTag(param.ID)
}

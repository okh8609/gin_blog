package dao

import "github.com/okh8609/gin_blog/internal/model"
import "github.com/okh8609/gin_blog/pkg/app"

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.BlogTag{Name: name, State: state}
	return tag.Count(d.gormDB)
}

func (d *Dao) GetTag(id uint32, state uint8) (model.BlogTag, error) {
	tag := model.BlogTag{BlogBasicInfo: &model.BlogBasicInfo{ID: id}, State: state}
	return tag.Get(d.gormDB)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.BlogTag, error) {
	tag := model.BlogTag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.gormDB, pageOffset, pageSize)
}

func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*model.BlogTag, error) {
	tag := model.BlogTag{State: state}
	return tag.ListByIDs(d.gormDB, ids)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.BlogTag{
		Name:  name,
		State: state,
		BlogBasicInfo: &model.BlogBasicInfo{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.gormDB)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.BlogTag{
		BlogBasicInfo: &model.BlogBasicInfo{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.gormDB, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.BlogTag{BlogBasicInfo: &model.BlogBasicInfo{ID: id}}
	return tag.Delete(d.gormDB)
}

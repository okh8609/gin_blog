package service

import (
	"context"

	// otgorm "github.com/okh8609/opentracing-gorm"

	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	// svc := Service{ctx: ctx}
	// svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))

	ss := Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
	return ss
}

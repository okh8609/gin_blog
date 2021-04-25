package service

import (
	"context"

	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	ss := Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
	return ss
}

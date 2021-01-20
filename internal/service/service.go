package service

import (
	"github.com/weirubo/blog-go/global"
	"github.com/weirubo/blog-go/internal/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() Service {
	s := Service{}
	// 全局变量
	s.dao = dao.New(global.DBEngine)
	return s
}

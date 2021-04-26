package service

import (
	"context"
	"errors"

	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/pkg/jwt"
	"github.com/okh8609/gin_blog/pkg/utils"
)

// Auth

func (s *Service) CreateAuth(param *CreateAuthParam) error {
	return s.dao.CreateAuth(param.UUID, utils.HashPassword(param.Password))
}

func (s *Service) VerifyAuth(param *VerifyAuthParam) (string, error) {
	//沒有 error 就核發 JWT token 否則回傳error跟空字串
	auth, err := s.dao.GetAuth(param.UUID)
	global.MyLogger.Debugf(context.Background(), "# Auth: %v", auth)
	if err != nil {
		return "", err
	}

	if auth.Password == utils.HashPassword(param.Password) {
		return jwt.GenerateJWTToken(param.UUID), nil
	}
	return "", errors.New("wrong password")
}

func (s *Service) UpdateAuth(param *UpdateAuthParam) error {
	auth, err := s.dao.GetAuth(param.UUID)
	if err != nil {
		return err
	}
	if auth.Password != utils.HashPassword(param.Password) {
		return errors.New("wrong password")
	}
	return s.dao.UpdateAuth(param.UUID, utils.HashPassword(param.NewPassword))
}

func (s *Service) DeleteAuth(param *DeleteAuthParam) error {
	auth, err := s.dao.GetAuth(param.UUID)
	if err != nil {
		return err
	}
	if auth.Password != utils.HashPassword(param.Password) {
		return errors.New("wrong password")
	}
	return s.dao.DeleteAuth(param.UUID)
}

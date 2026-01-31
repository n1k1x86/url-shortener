package service

import (
	"context"
	shortener_models "url-shortener/shortener/models"
	shortener_repo "url-shortener/shortener/repo"
)

type Service struct {
	repo *shortener_repo.Repo
}

func NewService(repo *shortener_repo.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetLinkByShort(ctx context.Context, short string, user_id int64) (string, error) {
	link, err := s.repo.GetLinkByShort(ctx, short, user_id)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (s *Service) GetAllLinks(ctx context.Context, user_id int64) ([]shortener_models.LinkRecord, error) {
	links, err := s.repo.GetAllLinks(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (s *Service) DeleteLink(ctx context.Context, short string, user_id int64) (bool, error) {
	ok, err := s.repo.DeleteLink(ctx, short, user_id)
	if err != nil {
		return ok, err
	}
	return ok, nil
}

func (s *Service) ShortLink(ctx context.Context, source string, short string, user_id int64) (bool, error) {
	ok, err := s.repo.ShortLink(ctx, source, short, user_id)
	if err != nil {
		return ok, err
	}
	return ok, nil
}

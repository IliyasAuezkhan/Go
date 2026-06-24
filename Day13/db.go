package main
import (
	"context"
)

type DBRepository interface {
	GetRoleByID(ctx context.Context, id int) (string, error)
}

type UserService struct {
	repo DBRepository
}

func NewUserService(r DBRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) IsAdmin(ctx context.Context, id int) bool {
	role, err := s.repo.GetRoleByID(ctx, id)
	if err != nil {
		return false
	}
	return role == "admin"
}
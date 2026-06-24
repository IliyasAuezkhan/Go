package main
import (
	"context"
	"errors"
	"testing"
)

type MockDB struct {
	FakeGetRole func(id int) (string, error)
}

func (m *MockDB) GetRoleByID(ctx context.Context, id int) (string, error) {
	return m.FakeGetRole(id)
}

func TestIsAdmin(t *testing.T) {
	mockSuccess := &MockDB{
		FakeGetRole: func (id int) (string, error) {
			return "admin", nil
		},
	}
	service1 := NewUserService(mockSuccess)
	res := service1.IsAdmin(context.Background(), 1)
	if res != true {
		t.Errorf("Ожидали true дляадмина но получили %v", res)
	}

	mockError := &MockDB{
		FakeGetRole: func(id int) (string, error) {
			return "", errors.New("database timeout")
		},
	}
	service2 := NewUserService(mockError)
	res = service2.IsAdmin(context.Background(), 2)
	if res != false {
		t.Errorf("Ожидали false при ошибке базы, но получили %v", res)
	}
}
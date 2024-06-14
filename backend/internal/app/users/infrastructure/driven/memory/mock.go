package memory

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/set"
)

type userMemory struct {
	set set.Set[string, *domain.User]
}

func NewUserMemory() userMemory {
	return userMemory{
		set: set.Set[string, *domain.User]{},
	}
}

func (u *userMemory) Save(_ context.Context, user *domain.User) error {
	ok := u.set.Exist(user.Name())
	if ok {
		return fmt.Errorf("%w = user with username '%s' already exists", domain.DuplicateUser, user.Name())
	}
	newRole, _ := domain.NewRole(user.Role().ID(), user.Role().Type())
	newUser, _ := domain.NewUser(user.ID(), user.Name(), user.Password(), newRole)
	u.set.Add(user.Name(), newUser)

	return nil
}

func (u *userMemory) Get(_ context.Context, username *domain.UserName) (*domain.User, uint8, error) {
	ok := u.set.Exist(username.Name())
	if !ok {
		return nil, 0, fmt.Errorf("%w = User with username '%s' not found", domain.UserNotFound, username.Name())
	}

	return u.set.GetByItem(username.Name()), 1, nil
}

package postgresql

import "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"

// BondSchema represnts a DTO(Data Transfer Object)
type UserSchema struct {
	ID          string
	Name        string
	Password    string
	Role        RoleSchema
	Permissions uint8
}

// RoleSchema represnts a DTO(Data Transfer Object)
type RoleSchema struct {
	ID   uint8
	Type string
}

func (r *RoleSchema) ToBusiness() (*domain.Role, error) {
	return domain.NewRole(r.ID, r.Type)
}

func (u *UserSchema) ToBusiness() (*domain.User, error) {
	role, err := u.Role.ToBusiness()
	if err != nil {
		return nil, err
	}

	return domain.NewUser(u.ID, u.Name, u.Password, role)
}

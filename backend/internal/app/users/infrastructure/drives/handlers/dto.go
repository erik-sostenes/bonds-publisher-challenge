package handlers

import "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"

type UserRequest struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Role     RoleRequest `json:"role"`
}

func (a *UserRequest) ToBusiness() (*domain.User, error) {
	role, err := a.Role.ToBusiness()
	if err != nil {
		return nil, err
	}

	return domain.NewUser(a.ID, a.Name, a.Password, role)
}

type RoleRequest struct {
	ID   uint8  `json:"id"`
	Type string `json:"type"`
}

func (r *RoleRequest) ToBusiness() (*domain.Role, error) {
	return domain.NewRole(r.ID, r.Type)
}

type RolesRequest []*RoleRequest

func (r *RolesRequest) ToBusiness() ([]*domain.Role, error) {
	roles := make([]*domain.Role, 0, len(*r))

	for _, role := range *r {
		rlBusiness, err := role.ToBusiness()
		if err != nil {
			return roles, err
		}
		roles = append(roles, rlBusiness)
	}

	return roles, nil
}

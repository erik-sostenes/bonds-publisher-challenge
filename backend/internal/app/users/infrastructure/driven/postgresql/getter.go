package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
)

type userGetter struct {
	conn *sql.DB
}

func NewUserGetter(conn *sql.DB) ports.UserGetter {
	if conn == nil {
		panic("missing sql.DB dependency")
	}
	return &userGetter{
		conn: conn,
	}
}

func (u *userGetter) Get(ctx context.Context, userId *domain.UserID) (user *domain.User, permissions uint8, err error) {
	const sqlQueryGetUser = `
		SELECT 
  			us.id,
  			us.name,
  			us.password,
  			ro.id,
  			ro.type,
  			SUM((permission::VARCHAR)::INTEGER) AS permissions
		FROM users us
  			INNER JOIN users_roles ur ON ur.user_id = us.id
			INNER JOIN roles ro ON ro.id = ur.role_id
			INNER JOIN permissions pe ON pe.role_id = ro.id
		WHERE us.id = $1
		GROUP BY us.id, ro.id`

	var userSchema UserSchema
	err = u.conn.QueryRowContext(ctx, sqlQueryGetUser, userId.ID()).Scan(
		&userSchema.ID,
		&userSchema.Name,
		&userSchema.Password,
		&userSchema.Role.ID,
		&userSchema.Role.Type,
		&userSchema.Permissions,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, fmt.Errorf("%w = User with id '%s' not found", domain.UserNotFound, userId.ID())
		}

		slog.ErrorContext(ctx, "database error", "msg", err.Error())
		return nil, 0, errors.New("an error occurred while retriever your account")

	}
	user, err = userSchema.ToBusiness()
	if err != nil {
		return nil, 0, err
	}

	return user, userSchema.Permissions, nil
}

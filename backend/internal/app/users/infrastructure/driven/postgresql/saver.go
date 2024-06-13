package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	"github.com/lib/pq"
)

type userSaver struct {
	conn *sql.DB
}

func NewUserSaver(conn *sql.DB) ports.UserSaver {
	if conn == nil {
		panic("missing sql.DB dependency")
	}
	return &userSaver{
		conn: conn,
	}
}

func (a *userSaver) Save(ctx context.Context, user *domain.User) error {
	tx, err := a.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	return transaction(tx, func() (err error) {
		const sqlQueryInsertUser = `
			INSERT INTO users(
				id,
				name,
				password
			) VALUES($1, $2, $3)`

		_, err = tx.ExecContext(ctx, sqlQueryInsertUser, user.ID(), user.Name(), user.Password())
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" { // 23505 is the PostgreSQL error code for unique violation
					return fmt.Errorf("%w = User with id '%s' already exists", domain.DuplicateUser, user.ID())
				}
			}

			slog.ErrorContext(ctx, "database error", "msg", err.Error())
			return errors.New("an error occurred while creating your account")
		}

		if err = a.insertUserRole(ctx, tx, user); err != nil {
			return
		}
		return
	})
}

func (a *userSaver) insertUserRole(ctx context.Context, tx *sql.Tx, user *domain.User) (err error) {
	const sqlQueryInsertUserRole = `INSERT INTO users_roles(user_id, role_id) VALUES($1, $2)`

	_, err = tx.ExecContext(ctx, sqlQueryInsertUserRole, user.ID(), user.Role().ID())
	if err != nil {
		slog.ErrorContext(ctx, "database error", "msg", err.Error())
		return errors.New("an error occurred while creating your account")
	}

	return
}

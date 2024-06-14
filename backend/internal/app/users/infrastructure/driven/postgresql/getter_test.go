package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/cmd/bootstrap/db"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
)

func Test_UserGetter(t *testing.T) {
	type FactoryFunc func() (ports.UserGetter, *sql.DB)

	const (
		sqlQueryInsertRole        = `INSERT INTO roles(id, type) VALUES($1, $2)`
		sqlQueryInsertPermissions = `INSERT INTO permissions(permission, role_id) VALUES($1, $2)`
		sqlQueryDeleteRole        = `DELETE FROM roles WHERE id = $1`
		sqlQueryDeletePermission  = `DELETE FROM permissions WHERE permission = $1`
		sqlQueryDeleteUser        = `DELETE FROM users WHERE id = $1`
		sqlQueryDeleteUserRole    = `DELETE FROM users_roles WHERE user_id = $1`
	)

	tdt := map[string]struct {
		user          UserSchema
		factoryFunc   FactoryFunc
		expectedError error
	}{
		"Given an existing valid user, it will be retrieved from postgresql": {
			user: UserSchema{
				ID:          "ba1dc545-90a0-4501-af99-8a5944ca38c4",
				Name:        "Erik Sostenes Simon",
				Password:    "password",
				Permissions: 4,
				Role: RoleSchema{
					ID:   1,
					Type: "USER",
				},
			},

			factoryFunc: func() (ports.UserGetter, *sql.DB) {
				conn := db.PostgreSQLInjector()
				userSchema := UserSchema{
					ID:          "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					Name:        "Erik Sostenes Simon",
					Password:    "password",
					Permissions: 4,
					Role: RoleSchema{
						ID:   1,
						Type: "USER",
					},
				}

				user, err := userSchema.ToBusiness()
				if err != nil {
					t.Fatal(err)
				}

				saver := NewUserSaver(conn)
				if err = saver.Save(context.TODO(), user); err != nil {
					t.Fatal(err)
				}
				return NewUserGetter(conn), conn
			},
		},
		"Given a non-existing valid user, it will not be retrieved from postgresql": {
			user: UserSchema{
				ID:          "ba1dc545-90a0-4501-af99-8a5944ca38c4",
				Name:        "Erik Sostenes Simon",
				Password:    "password",
				Permissions: 4,
				Role: RoleSchema{
					ID:   1,
					Type: "USER",
				},
			},

			factoryFunc: func() (ports.UserGetter, *sql.DB) {
				conn := db.PostgreSQLInjector()
				return NewUserGetter(conn), conn
			},
			expectedError: domain.UserNotFound,
		},
	}

	ctx := context.Background()

	// setUp
	conn := db.PostgreSQLInjector()
	if err := func() (err error) {
		if _, err = conn.ExecContext(ctx, sqlQueryInsertRole, 1, "USER"); err != nil {
			return
		}
		if _, err = conn.ExecContext(ctx, sqlQueryInsertPermissions, `1`, 1); err != nil {
			return
		}
		return
	}(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if _, err := conn.ExecContext(ctx, sqlQueryDeletePermission, `1`); err != nil {
			t.Fatal(err)
		}
		if _, err := conn.ExecContext(ctx, sqlQueryDeleteRole, 1); err != nil {
			t.Fatal(err)
		}
	})

	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T) {
			getter, conn := tsc.factoryFunc()
			userSchema, err := tsc.user.ToBusiness()
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := conn.ExecContext(ctx, sqlQueryDeleteUserRole, userSchema.ID()); err != nil {
					t.Fatal(err)
				}
				if _, err := conn.ExecContext(ctx, sqlQueryDeleteUser, userSchema.ID()); err != nil {
					t.Fatal(err)
				}
			})

			username := domain.UserName(userSchema.Name())
			_, _, err = getter.Get(ctx, &username)
			asUser := domain.UserError(0)

			if errors.As(err, &asUser) {
				if !errors.Is(asUser, tsc.expectedError) {
					t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asUser)
				}
			} else if err != nil {
				t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, err)
			}
		})
	}
}

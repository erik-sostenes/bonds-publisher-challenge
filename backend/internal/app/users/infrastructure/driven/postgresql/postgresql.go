package postgresql

import (
	"database/sql"
	"fmt"
)

// transaction executes a function within the context of a database transaction
func transaction(tx *sql.Tx, f func() error) error {
	if err := f(); err != nil {
		_ = tx.Rollback()

		return fmt.Errorf("f %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit %w", err)
	}

	return nil
}

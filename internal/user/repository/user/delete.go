package user

import (
	"context"
	"fmt"
)

func (r *Repo) Delete(ctx context.Context, userID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", userTable)

	if _, err := r.DB.ExecContext(ctx, query, userID); err != nil {
		return fmt.Errorf("can not delete:%v", err)
	}
	return nil
}

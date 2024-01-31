package user

import (
	"context"
	"fmt"
	"github.com/vshigimoto/LinkedIn-clone/internal/user/entity"
)

func (r *Repo) Update(ctx context.Context, user entity.User) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, surname = $2, about = $3 WHERE id = $4", userTable)

	if _, err := r.DB.ExecContext(ctx, query, user.Name, user.Surname, user.About, user.Id); err != nil {
		return fmt.Errorf("can not update:%v", err)
	}
	return nil
}

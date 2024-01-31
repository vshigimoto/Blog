package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/vshigimoto/LinkedIn-clone/internal/user/entity"
)

func (r *Repo) Get(ctx context.Context, userID int) (entity.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)

	var user entity.User
	if err := r.DB.QueryRowContext(ctx, query, userID).
		Scan(&user.Id, &user.Name, &user.Surname, &user.About); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, nil
		}

		return entity.User{}, fmt.Errorf("can not get:%v", err)
	}
	return user, nil
}

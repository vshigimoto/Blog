package user

import (
	"context"
	"fmt"
	"github.com/vshigimoto/LinkedIn-clone/internal/user/entity"
)

func (r *Repo) Create(ctx context.Context, user entity.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, surname, about) VALUES ($1, $2, $3) RETURNING id", userTable)

	var userID int
	if err := r.DB.QueryRowContext(ctx, query, "param2", "param1").Scan(&userID); err != nil {
		return 0, fmt.Errorf("can't create user: %v", err)
	}
	return userID, nil
}

package user

import (
	"context"
	"fmt"
	entity "github.com/vshigimoto/Blog/internal/blog/entity/user"
)

func (r *Repo) Create(ctx context.Context, user entity.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (last_name, first_name) VALUES ($1, $2) RETURNING id", userTable)

	var userID int
	if err := r.DB.QueryRowxContext(ctx, query, "param2", "param1").Scan(&userID); err != nil {
		return 0, fmt.Errorf("can't create user: %v", err)
	}
	return userID, nil
}

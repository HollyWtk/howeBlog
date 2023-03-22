package repo

import (
	"context"
	"github.com/blog-admin/internal/data"
)

type UserRepo interface {
	FindMember(ctx context.Context, username string, pwd string) (user *data.User, err error)
}

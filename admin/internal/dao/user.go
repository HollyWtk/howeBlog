package dao

import (
	"context"
	"github.com/blog-admin/internal/data"
	"github.com/blog-admin/internal/database/gorms"
	"gorm.io/gorm"
)

type UserDao struct {
	conn *gorms.GormConn
}

func (u *UserDao) FindMember(ctx context.Context, username string, pwd string) (*data.User, error) {
	var user *data.User
	err := u.conn.Session(ctx).Where("username=? and password=?", username, pwd).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}
func NewUserDao() *UserDao {
	return &UserDao{
		conn: gorms.New(),
	}
}

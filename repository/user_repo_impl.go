package repository

import (
	"context"
	"errors"
	"github.com/example/db"
	"github.com/example/model"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.SQL
}

func NewUserRepo(sql *db.SQL) UserRepo {
	return UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {

	statement := `
		INSERT INTO users(user_id,full_name,role,created_at,password,email)
		VALUES (:user_id,:full_name,:role,:created_at,:password,:email)
	`
	now := time.Now()
	user.CreatedAt = now
	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errors.New("Ng dung ton tai")
			}
		}
		return user, errors.New("dang ky that bai")
	}
	return user, nil
}

func (u UserRepoImpl) SelectUserByEmail(context context.Context, email string) (model.User, error) {
	return model.User{}, nil
}

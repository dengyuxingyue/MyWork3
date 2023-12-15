package serializer

import "work/model"

type User struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"creat_at" form:"creat_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}

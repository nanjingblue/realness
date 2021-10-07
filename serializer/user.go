package serializer

import (
	"realness/model"
	"time"
)

// User 用户序列化器
type User struct {
	ID        int32   `json:"id"`
	UserName  string `json:"user_name"`
	Email string	`json:"email"`
	Role model.UserRole	`json:"role"`
	Avatar    string `json:"avatar"`
	Gender int `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Sign string `json:"sign"`
	CreatedAt int64  `json:"created_at"`
}

type UserLike struct {
	ID		uint `json:"id"`
	UserID 	int32 `json:"user_id"`
	PostID		int32 `json:"post_id"`
	UpdatedAt int64 `json:"updated_at"`
	CreatedAt int64 `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Email: user.Email,
		Role: user.Role,
		Avatar: user.Avatar,
		Gender: user.Gender,
		Birthday: user.Birthday,
		Sign: user.Sign,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

func BuildUserLike(u model.UserLike) UserLike {
	return UserLike{
		ID: u.ID,
		UserID: u.UserID ,
		PostID: u.PostID,
		UpdatedAt: u.UpdatedAt.Unix(),
		CreatedAt: u.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

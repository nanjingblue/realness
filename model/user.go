package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserRole int32

const (
	BannedUserRole   UserRole = -100
	SuperUserRole             = 0
	AdminRole                 = 1
	DeleterRole               = 2
	UnDeleterRole             = 3
	Deleter2Role              = 20
	Deleter3Role              = 21
	NormalUserRole            = 50
	UnregisteredRole          = 100
)

// User 用户模型
type User struct {
	ID             int32  `gorm:"primaryKey;autoIncrement;not null"`
	UserName       string		`gorm:"type:varchar(20);not null"`
	PasswordDigest string
	Email    string    `gorm:"varchar(20);not null;index"`
	Role		UserRole
	Avatar  string 	`gorm:"size:1000;default:'https://z3.ax1x.com/2021/06/24/RQIySs.png'"`
	Gender   int       `gorm:"default:0"`
	Birthday time.Time `gorm:"default:'1970-01-01'"`
	Sign     string    `gorm:"varchar(50);default:'这个人很懒，什么都没有留下'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

type UserLike struct {
	gorm.Model
	UserID int32
	PostID int32
}
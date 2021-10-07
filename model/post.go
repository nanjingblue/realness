package model

import (
	"gorm.io/gorm"
	"time"
)

type ReportType string

const (
	UserReport        ReportType = "UserReport"
	UserReportFold    ReportType = "UserReportFold"
	UserDelete        ReportType = "UserDelete" // delete, no ban
	AdminTag          ReportType = "AdminTag"
	AdminDeleteAndBan ReportType = "AdminDeleteBan" // delete, ban
	AdminUndelete     ReportType = "Undelete"       // undelete + unban
	AdminUnban        ReportType = "AdminUnban"     // delete + unban
	//	For now, there's no "undelete + no unban" option
)

type Post struct {
	ID             int32  `gorm:"primaryKey;autoIncrement;not null"`
	UserID     int32
	Text         string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram;type: varchar(10000) NOT NULL"`
	Cover		string
	Tag			string `gorm:"index;type:varchar(60) NOT NULL"`
	Type 		string `gorm:"type:varchar(20) NOT NULL"`
	LikeNum      int32  `gorm:"index"`
	ReplyNum     int32  `gorm:"index"`
	ReportNum    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Comment struct {
	ID      int32 `gorm:"primaryKey;autoIncrement;not null"`
	ReplyTo int32 `gorm:"index"`
	//Post         Post
	PostID int32 `gorm:"index"`
	//User         User
	UserID       int32
	Text         string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram;type: varchar(10000) NOT NULL"`
	Tag          string `gorm:"index;type:varchar(60) NOT NULL"`
	Type         string `gorm:"type:varchar(20) NOT NULL"`
	Name         string `gorm:"type:varchar(60) NOT NULL"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Report struct {
	ID int32 `gorm:"primaryKey;autoIncrement;not null"`
	//User           User
	UserID int32
	//ReportedUser   User
	ReportedUserID int32
	//Post           Post
	PostID int32
	//Comment        Comment
	CommentID int32
	Reason    string     `gorm:"type: varchar(1000) NOT NULL"`
	Type      ReportType `gorm:"type:varchar(20) NOT NULL"`
	IsComment bool
	Weight    int32
	CreatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

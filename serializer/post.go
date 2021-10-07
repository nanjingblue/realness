package serializer

import (
	"realness/model"
)

// Post 串序列化器
type Post struct {
	ID        int32   `json:"id"`
	UserID		int32 `json:"user_id"`
	Text         string `json:"text"`
	Cover		string	`json:"cover"`
	Tag			string `json:"tag"`
	Type 		string `json:"type"`
	LikeNum      int32  `json:"like_num"`
	ReplyNum     int32  `json:"reply_num"`
	ReportNum    int32 `json:"report_num"`
	CreatedAt int64  `json:"created_at"`
}

// BuildPost 序列化串
func BuildPost(post model.Post) Post {
	return Post {
		ID: post.ID,
		UserID: post.UserID,
		Text: post.Text,
		Cover: post.Cover,
		Tag: post.Tag,
		Type: post.Type,
		LikeNum: post.LikeNum,
		ReplyNum: post.ReplyNum,
		ReportNum: post.ReportNum,
		CreatedAt: post.CreatedAt.Unix(),
	}
}

// BuildPostResponse 序列化用户响应
func BuildPostResponse(post model.Post) Response {
	return Response{
		Data: BuildPost(post),
	}
}

// BuildPosts 序列化所有的串
func BuildPosts(items []model.Post) (posts []Post)  {
	for _, item := range items {
		post := BuildPost(item)
		posts = append(posts, post)
	}
	return posts
}

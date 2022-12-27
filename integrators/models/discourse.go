package models

import (
	"fmt"
	"time"
)

type SearchResult struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	PostNumber int64     `json:"post_number"`
	LikeCount  int64     `json:"like_count"`
	Blurb      string    `json:"blurb"`
	TopicID    int64     `json:"topic_id"`
}

type PostResult struct {
	PostStream struct {
		UserPosts []UserPost `json:"posts"`
	} `json:"post_stream"`
}

type UserPost struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Cooked       string    `json:"cooked"`
	PostType     int64     `json:"post_type"`
	Reads        int64     `json:"reads"`
	ReadersCount int64     `json:"readers_count"`
}

type DiscoursePostData struct {
	ID           int64
	TopicID      int64
	PostNumber   int64
	Reads        int64
	ReadersCount int64
	Username     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LikeCount    int64
	Type         string
	Language     string
	Description  string
	Metadata     map[string]string
}

// ToDiscourseDataModel creates data record using the fetched data
func ToDiscourseDataModel(post Post, userPost UserPost) *DiscoursePostData {
	return &DiscoursePostData{
		ID:           post.ID,
		TopicID:      post.TopicID,
		PostNumber:   post.PostNumber,
		Reads:        userPost.Reads,
		ReadersCount: userPost.ReadersCount,
		Username:     userPost.Username,
		CreatedAt:    userPost.CreatedAt,
		UpdatedAt:    userPost.UpdatedAt,
		LikeCount:    post.LikeCount,
		Language:     "en", // assuming all Discourse posts are in English
		Description:  userPost.Cooked,
		Metadata: map[string]string{
			"post_type": fmt.Sprintf("%d", userPost.PostType),
		},
	}
}

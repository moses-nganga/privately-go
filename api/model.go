package api

import (
	"time"
)

type Knot struct{
	ID   string `json:"id"`
	Name string `json:"name"`
	CoverImage string `json:"cover_image"`
	IsDefaultKnot	bool `json:"is_default_knot"`
	CreatedBy string `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
type Album struct {
	ID string `json:"id"`
	Name string `json:"name"`
	KnotId string `json:"knot"`
	CoverImage string `json:"cover_image"`
	IsTimelineAlbum bool `json:"is_timeline_album"`
	CreatedBy string `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
type Moment struct {
	ID string `json:"id"`
	AlbumId string `json:"album_id"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	Likes int `json:"likes"`
	CreatedBy string `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
type MomentLike struct {
	ID string `json:"id"`
	MomentId string `json:"moment_id"`
	LikedBy string `json:"liked_by"`
	LikedAt time.Time `json:"liked_at"`
}
type MomentComment struct {
	ID string `json:"id"`
	MomentId string `json:"moment_id"`
	Comment string `json:"comment"`
	CommentAt string `json:"comment_at"`
	CommentBy string `json:"comment_by"`
	CommentParent string `json:"comment_parent"`
}
type Feed struct {
	ID string `json:"id"`
	Actor string `json:"actor"`
	Verb string `json:"verb"`
	Moments []Moment `json:"moments"`
	CreatedAt string `json:"created_at"`
} 
type Notification struct {
	ID string `json:"id"`
	NotificationText string `json:"notification_text"`
	NotificationType string `json:"notification_type"`
	NotificationTo	string `json:"notification_to"`
	CreatedAt time.Time `json:"created_at"`
}
type KnotMember struct {
	ID string `json:"id"`
	KnotId string `json:"knot_id"`
	UserId string `json:"user_id"`
	AddedAt time.Time `json:"added_at"`
}
type User struct {
	ID string `json:"id"`
	FullName string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	DefaultKnot	string `json:"default_knot"`
	CreatedAt time.Time `json:"created_at"`
}

type Notifications []Notification
type Users []User
type NewsFeed []Feed
type Moments []Moment
type Comments []MomentComment
type Likes []MomentLike
type Albums []Album
type Knots []Knot 

package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Repository struct {
}

const DB_SERVER = ""
const DB_NAME = ""
const DB_USER = ""
const DB_PASSWORD = ""
const KNOT_TBL = "pl_knot"
const ALBUM_TBL = "pl_album"
const MOMENT_TBL = "pl_moment"
const LIKE_TBL = "pl_like"
const COMMENT_TBL = "pl_comment"
const FEED_TBL = "pl_feed"
const NOTIFICATION_TBL = "pl_notification"
const USER_TBL = "pl_user"

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_SERVER+":3306)/"+DB_NAME+"?parseTime=true&loc=US%2FPacific")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func (r Repository) newKnot(knot Knot) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + KNOT_TBL + " (id,name,cover_image,is_default_knot,created_by,created_at) VALUES(?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(knot.ID, knot.Name, knot.CoverImage, knot.IsDefaultKnot, knot.CreatedBy, knot.CreatedAt)

	defer db.Close()

	fmt.Println("Added a new Knot")

	return true
}
func (r Repository) newAlbum(album Album) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + ALBUM_TBL + " (id,name,knot_id,cover_image,is_timeline_album,created_by,created_at) VALUES(?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(album.ID, album.Name, album.KnotId, album.CoverImage, album.IsTimelineAlbum, album.CreatedBy, album.CreatedAt)

	defer db.Close()

	fmt.Println("Added an Album")

	return true
}
func (r Repository) newMoment(moment Moment) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + MOMENT_TBL + " (id,album_id,caption,photo_url,likes,created_by,created_at) VALUES(?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(moment.ID, moment.AlbumId, moment.Caption, moment.PhotoUrl, moment.Likes, moment.CreatedBy, moment.CreatedAt)

	defer db.Close()

	fmt.Println("added a Moment")

	return true
}
func (r Repository) newLike(like MomentLike) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + LIKE_TBL + " (id,moment_id,liked_by,liked_at) VALUES(?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(like.ID, like.MomentId, like.LikedBy, like.LikedAt)

	defer db.Close()

	fmt.Println("Added a Like")

	return true
}
func (r Repository) newComment(comment MomentComment) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + COMMENT_TBL + " (id,moment_id,comment,comment_at,comment_by,comment_parent) VALUES(?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(comment.ID, comment.MomentId, comment.Comment, comment.CommentAt, comment.CommentBy, comment.CommentParent)

	defer db.Close()

	fmt.Println("Added a Comment")

	return true
}
func (r Repository) newFeed(feed Feed) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + FEED_TBL + " (id,actor,verb,moments,created_at) VALUES(?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(feed.ID, feed.Actor, feed.Verb, feed.Moments, feed.CreatedAt)

	defer db.Close()

	fmt.Println("Added to Feed")

	return true
}
func (r Repository) newNotification(notification Notification) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + NOTIFICATION_TBL + " (id,notificationText,notificationType,notificationTo,createdAt) VALUES(?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(notification.ID, notification.NotificationText, notification.NotificationType, notification.NotificationTo, notification.CreatedAt)

	defer db.Close()

	fmt.Println("Added a Notification")

	return true
}

func (r Repository) getKnots(userId string) Knots{
	results := Knots{}

	db := dbConnect()

}

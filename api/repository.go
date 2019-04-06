package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Repository struct {
}

const DB_SERVER = "fnx6frzmhxw45qcb.cbetxkdyhwsb.us-east-1.rds.amazonaws.com"
const DB_NAME = "mq2nzp21vayc85l6"
const DB_USER = "grv12hh0j4qbi6oo"
const DB_PASSWORD = "zc6zojohv5019wa5"
const KNOT_TBL = "pl_knot"
const KNOT_MEMBERS_TBL="pl_knot_members"
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
func (r Repository) newUser(user User) bool {
	db := dbConnect()

	stmt, err := db.Prepare("INSERT INTO " + USER_TBL + " (id,full_name,default_knot,created_at) VALUES(?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(user.ID, user.FullName,user.DefaultKnot,user.CreatedAt)

	defer db.Close()

	fmt.Println("Added a new User")

	return true
}
func (r Repository) UpdateProfilePhoto(user User) bool {

	db := dbConnect()

	stmt, err := db.Prepare("UPDATE " + USER_TBL + " SET profile_photo=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(user.ProfilePhoto,user.ID)

	defer db.Close()

	fmt.Println("Updated Profile Photo")
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

	stmt, err := db.Prepare("INSERT INTO " + NOTIFICATION_TBL + " (id,notification_text,notification_type,notification_to,created_at) VALUES(?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(notification.ID, notification.NotificationText, notification.NotificationType, notification.NotificationTo, notification.CreatedAt)

	defer db.Close()

	fmt.Println("Added a Notification")

	return true
}
func (r Repository) joinKnot(member KnotMember) bool{
	db := dbConnect()

	stmt,err := db.Prepare("INSERT INTO "+KNOT_MEMBERS_TBL +" (id,knot_id,user_id,added_at) VALUES(?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(member.ID,member.KnotId,member.UserId,member.AddedAt)

	defer db.Close()
	fmt.Println("Joined a group")

	return true
}
func (r Repository) getKnots(userId string) Knots{
	results := Knots{}

	db := dbConnect()

	stmt,err := db.Query("SELECT k.id,k.name,k.cover_image,k.is_default,k.created_at,k.created_by FROM "+KNOT_TBL+" k INNER JOIN "+KNOT_MEMBERS_TBL+" km ON k.id=km.knot_id WHERE created_by=?",userId)
	if err != nil {
		panic(err.Error())
	}
	for stmt.Next(){
		var knot Knot
		err := stmt.Scan(&knot.ID,knot.Name,knot.CoverImage,knot.IsDefaultKnot,knot.CreatedAt,knot.CreatedBy)
		if err != nil {
			panic(err.Error())
		}
		results = append(results,knot)
	}

	defer db.Close()
	
	return results
}
func (r Repository) getAlbums(knotId string) Albums{
	results := Albums{}

	db := dbConnect()

	stmt,err := db.Query("SELECT id,name,cover_image,is_timeline_album,created_by,created_at FROM "+ALBUM_TBL+" WHERE knot_id=?",knotId)
	if err!= nil{
		panic(err.Error())
	}
	for stmt.Next() {
		var album Album
		err := stmt.Scan(&album.ID,&album.Name,&album.CoverImage,&album.IsTimelineAlbum,&album.CreatedBy,&album.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		results = append(results,album)
	}

	defer db.Close()

	return results
}
func (r Repository) getMoments(albumId string) Moments{
	results := Moments{}

	db := dbConnect()

	stmt,err := db.Query("SELECT id,caption,photo_url,likes,created_by,created_at FROM "+MOMENT_TBL+" WHERE album_id=?",albumId)
	if err != nil {
		panic(err.Error())
	}
	for stmt.Next(){
		var moment Moment
		if err != nil {
			panic(err.Error())
		}
		results = append(results,moment)
	}
	defer db.Close()

	return results
}
func (r Repository) getLikes(momentId string) Users{
	results := Users{}

	db := dbConnect()

	stmt,err := db.Query("SELECT u.fullname FROM "+LIKE_TBL+" l INNER JOIN "+USER_TBL+" u ON l.liked_by=a.id WHERE moment_id=?",momentId)
	if err != nil {
		panic(err.Error())
	}
	for stmt.Next() {
		var user User
		results = append(results, user)
	}
	defer db.Close()

	return results
}
func (r Repository) getComments(momentId string) Comments{
	results := Comments{}

	db := dbConnect()

	stmt,err := db.Query("SELECT c.id,c.comment,c.comment_at,u.fullname FROM "+COMMENT_TBL+" c INNER JOIN "+USER_TBL+" u ON c.comment_by=u.id WHERE moment_id=?",momentId)
	if err != nil {
		panic(err.Error())
	}
	for stmt.Next()  {
		var comment MomentComment
		results = append(results,comment)
	}
	defer db.Close()

	return results

}
func (r Repository) getNotifications(userId string)  Notifications{
	results := Notifications{}

	db := dbConnect()

	stmt,err := db.Query("SELECT id,notification_text,notification_type,created_at FROM "+NOTIFICATION_TBL+" WHERE notification_to=?",userId)
	if err != nil {
		panic(err.Error())
	}

	for stmt.Next(){
		var notification Notification
		results = append(results,notification)
	}
	defer db.Close()

	return results
}
func (r Repository) GetKnotMembers(knotId string) Users{
	results := Users{}

	db := dbConnect()

	stmt,err := db.Query("SELECT u.id,u.fullname,u.profilePhoto,u.default_knot FROM "+KNOT_MEMBERS_TBL+" k INNER JOIN "+USER_TBL+" u ON k.user_id=u.id WHERE k.knot_id=?",knotId)

	if err != nil {
		panic(err.Error())
	}

	for stmt.Next(){
		var user User
		results = append(results,user)
	}
	defer db.Close()

	return results
}
package api

import (
	"encoding/json"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	gContext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	)
var user User

const bearer  = "Bearer"

type Controller struct{
	repository Repository
}

func initializeApp() *firebase.App  {
	opt := option.WithCredentialsFile("privately-22fd2-firebase-adminsdk-zd86p-0a2117baee.json")
	app,err := firebase.NewApp(context.Background(),nil,opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n",err)
	}
	return app
}
func getUserId(r *http.Request) string {
	//Get the user object from the Request
	user := gContext.Get(r,"decoded").(User)
	//Extract the ID and return it
	return user.ID
}
//Verify Token
func verifyIDToken(ctx context.Context,app *firebase.App, idToken string) (*auth.Token,error){
	client,err := app.Auth(context.Background())
	if err != nil {
		return nil,err
	}
	log.Printf("error getting Auth client: %v\n",err)

	token, err := client.VerifyIDToken(ctx,idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n",err)
		return nil,err
	}
	log.Printf("Id token Verified")

	return token,nil
}
func Authenticate(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		//Get Token from the Request headers
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken,"Bearer ")
			reqToken = splitToken[1]

			if len(reqToken) > 0 {
				//Initialize SDK
				app := initializeApp()

				token,err := verifyIDToken(r.Context(),app,reqToken)
				if err != nil {
					json.NewEncoder(w).Encode("Authorization Failed")
				}else {
					user.ID = token.UID
					gContext.Set(r,"decoded",user)
					next(w,r)
				}
			}else {
				json.NewEncoder(w).Encode("Invalid Token")
			}
		}else {
			json.NewEncoder(w).Encode("An authorization Header is required")
		}
	});
}
func (c *Controller) NewKnot(w http.ResponseWriter,r *http.Request) {
	var knot Knot
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Knot",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Knot",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &knot); err != nil {
			w.WriteHeader(422)
			log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	knot.CreatedBy=user.ID
	knot.CreatedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newKnot(knot)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewAlbum(w http.ResponseWriter,r *http.Request) {
	var album Album
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Album",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Album",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &album); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	album.CreatedBy=user.ID
	album.CreatedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newAlbum(album)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewMoment(w http.ResponseWriter,r *http.Request) {
	var moment Moment
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Moment",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Moment",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &moment); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	moment.CreatedBy=user.ID
	moment.CreatedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newMoment(moment)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewLike(w http.ResponseWriter,r *http.Request) {
	var like MomentLike
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Like",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Like",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &like); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	like.LikedBy=user.ID
	like.LikedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newLike(like)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewComment(w http.ResponseWriter,r *http.Request) {
	var comment MomentComment
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Comment",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Comment",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &comment); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	comment.CommentBy=user.ID
	comment.CommentAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newComment(comment)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewFeed(w http.ResponseWriter,r *http.Request) {
	var feed Feed
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Feed",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Feed",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &feed); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	feed.CreatedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newFeed(feed)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewNotification(w http.ResponseWriter,r *http.Request) {
	var notification Notification
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New Notification",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New Notification",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &notification); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	notification.NotificationTo=user.ID
	notification.CreatedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newNotification(notification)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) JoinKnot(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	KnotId := vars["id"]

	var member KnotMember

	member.UserId=user.ID
	member.KnotId=KnotId
	member.AddedAt=time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.joinKnot(member)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) NewUser(w http.ResponseWriter,r *http.Request) {
	var newUser User
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Creating New User",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Creating New User",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//update this users Id and the created time to now
	newUser.ID=user.ID
	newUser.CreatedAt = time.Now()

	//The unMarshal worked,write to the DB
	success := c.repository.newUser(user)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) UpdateProfilePhoto(w http.ResponseWriter,r *http.Request) {
	var newUser User
	//Read the body of the request
	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	//Handle the error
	if err != nil {
		log.Fatalln("Error Updating Profile Photo",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Close the response body
	if err := r.Body.Close(); err != nil{
		log.Fatalln("Error Uploading Profile Photo",err)
	}
	//Unmarshal the json into our Knots
	if err := json.Unmarshal(body, &newUser); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err);err!= nil {
			log.Fatalln("Error unmarshalling data",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	//The unMarshal worked,write to the DB
	success := c.repository.UpdateProfilePhoto(newUser)

	//Handle the error
	if !success{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set headers and update the response
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
func (c *Controller) GetKnots(w http.ResponseWriter,r *http.Request)  {
	//Get this users knots
	knots := c.repository.getKnots(user.ID)
	//Marshal then into JSON
	data,_:=json.Marshal(knots)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
func (c *Controller) GetKnotMembers(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	knotMemberId := vars["id"]
	//Get this users knots
	members := c.repository.GetKnotMembers(knotMemberId)
	//Marshal then into JSON
	data,_:=json.Marshal(members)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
func (c *Controller) GetAlbums(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	knotId := vars["id"]
	//Get this knots albums
	albums := c.repository.getAlbums(knotId)
	//Marshal then into JSON
	data,_:=json.Marshal(albums)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
func (c *Controller) GetMoments(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	albumId := vars["id"]
	//Get this albums moments
	moments := c.repository.getMoments(albumId)
	//Marshal then into JSON
	data,_:=json.Marshal(moments)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
func (c *Controller) GetLikes(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	momentId := vars["id"]
	//Get this moments likes
	likes := c.repository.getLikes(momentId)
	//Marshal then into JSON
	data,_:=json.Marshal(likes)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
func (c *Controller) GetComments(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	momentId := vars["id"]
	//Get this moments comments
	comments := c.repository.getComments(momentId)
	//Marshal then into JSON
	data,_:=json.Marshal(comments)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
func (c *Controller) GetNotifications(w http.ResponseWriter,r *http.Request)  {
	//Get this users knots
	notifications := c.repository.getNotifications(user.ID)
	//Marshal then into JSON
	data,_:=json.Marshal(notifications)
	//Setup the Response
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
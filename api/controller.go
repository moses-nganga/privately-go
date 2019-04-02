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
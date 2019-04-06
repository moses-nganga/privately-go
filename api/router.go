package api

import (
	"github.com/gorilla/mux"
	"net/http"
	)

var controller = &Controller{repository:Repository{}}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"SetUp",
		"POST",
		"/v0/setup",
		Authenticate(controller.NewUser),
	},
	Route{
		"JoinKnot",
		"GET",
		"/v0/knot/join/{id}",
		Authenticate(controller.JoinKnot),
	},
	Route{
		"NewKnot",
		"POST",
		"/v0/Knot/new",
		Authenticate(controller.NewKnot),
	},
	Route{
		"NewAlbum",
		"POST",
		"/v0/album/new",
		Authenticate(controller.NewAlbum),
	},
	Route{
		"NewMoment",
		"POST",
		"/v0/moment/new",
		Authenticate(controller.NewMoment),
	},
	Route{
		"NewLike",
		"POST",
		"/v0/like/new",
		Authenticate(controller.NewLike),
	},
	Route{
		"NewComment",
		"POST",
		"/v0/comment/new",
		Authenticate(controller.NewComment),
	},
	Route{
		"NewFeed",
		"POST",
		"/v0/feed/new",
		Authenticate(controller.NewFeed),
	},
	Route{
		"NewNotification",
		"POST",
		"/v0/notification/new",
		Authenticate(controller.NewNotification),
	},
	Route{
		"NewUser",
		"POST",
		"/v0/user/new",
		Authenticate(controller.NewUser),
	},
	Route{
		"UpdateProfilePhoto",
		"POST",
		"/v0/update-photo",
		Authenticate(controller.UpdateProfilePhoto),
	},
	Route{
		"GetKnots",
		"GET",
		"/v0/get-knots",
		Authenticate(controller.GetKnots),
	},
	Route{
		"GetKnotMembers",
		"GET",
		"v0/get-knot-members/{id}",
		Authenticate(controller.GetKnotMembers),
	},
	Route{
		"GetAlbums",
		"GET",
		"v0/get-albums/{id}",
		Authenticate(controller.GetAlbums),
	},
	Route{
		"GetMoments",
		"GET",
		"v0/get-moments/{id}",
		Authenticate(controller.GetMoments),
	},
	Route{
		"GetLikes",
		"GET",
		"v0/get-likes/{id}",
		Authenticate(controller.GetLikes),
	},
	Route{
		"GetComments",
		"GET",
		"v0/get-comments/{id}",
		Authenticate(controller.GetComments),
	},
	Route{
		"GetNotifications",
		"GET",
		"v0/get-notifications",
		Authenticate(controller.GetNotifications),
	},
}

func NewRouter() *mux.Router  {
	//Create a new router
	router := mux.NewRouter().StrictSlash(true)
	//Load it up
	for _,route :=range routes{
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

package privately

import (
	"github.com/gorilla/handlers"
	"privately/api"
	"net/http"
	"log"
	"os"
)

func main()  {
	//Get the port
	port := os.Getenv("PORT")

	if port=="" {
		log.Fatal("$port must be set")
	}

	//Initialize the new routes
	router := api.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods:= handlers.AllowedMethods([]string{"GET","POST","DELETE","PUT"})
	log.Fatal(http.ListenAndServe(":"+port,handlers.CORS(allowedOrigins,allowedMethods)(router)))
}

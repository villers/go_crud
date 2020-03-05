package tweet

import (
	"go_crud/pkg/container"
)

func LoadRoutes(app *container.Container) {
	app.Router.HandleFunc("/tweets", GetTweets(app)).Methods("GET")
	app.Router.HandleFunc("/tweets/{id:[0-9]+}", GetTweet(app)).Methods("GET")
	app.Router.HandleFunc("/tweets", PostTweet(app)).Methods("POST")
	app.Router.HandleFunc("/tweets/{id:[0-9]+}", UpdateTweet(app)).Methods("PUT")
}

package member

import (
	"go_crud/pkg/container"
)

func LoadRoutes(app *container.Container) {
	app.Router.HandleFunc("/members", GetMembers(app)).Methods("GET")
	app.Router.HandleFunc("/members/{id:[0-9]+}", GetMember(app)).Methods("GET")
	app.Router.HandleFunc("/members", PostMember(app)).Methods("POST")
	app.Router.HandleFunc("/members/{id:[0-9]+}", UpdateMember(app)).Methods("PUT")
}

package member

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_crud/pkg"
	"go_crud/pkg/container"
	"go_crud/pkg/member/models"
	"net/http"
)

func GetMembers(app *container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var members []models.Members
		app.DB.Find(&members)

		pkg.RespondWithJSON(w, http.StatusOK, members)
	}
}

func GetMember(app *container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var members models.Members

		if app.DB.First(&members, vars["id"]).RecordNotFound() {
			pkg.RespondWithError(w, http.StatusNotFound, "Not Found")
			return
		}

		pkg.RespondWithJSON(w, http.StatusOK, members)
	}
}

func PostMember(app *container.Container) http.HandlerFunc {

	type jsonQuery struct {
		Mail     string `json:"mail"`
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Passw    string `json:"passw"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var member jsonQuery
		err := json.NewDecoder(r.Body).Decode(&member)

		if err != nil {
			pkg.RespondWithError(w, http.StatusBadRequest, "Bad Request "+err.Error())
			return
		}

		app.DB.Create(&models.Members{
			Mail:       member.Mail,
			Fullname:   member.Fullname,
			Username:   member.Username,
			Passw:      member.Passw,
			EtatCompte: "0",
			LightMode:  "on",
		})

		pkg.RespondWithNoContent(w)
	}
}

func UpdateMember(app *container.Container) http.HandlerFunc {

	type jsonQuery struct {
		Mail     string `json:"mail"`
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Passw    string `json:"passw"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var memberQuery jsonQuery
		err := json.NewDecoder(r.Body).Decode(&memberQuery)

		if err != nil {
			pkg.RespondWithError(w, http.StatusBadRequest, "Bad Request "+err.Error())
			return
		}

		vars := mux.Vars(r)

		var member models.Members
		if app.DB.First(&member, vars["id"]).RecordNotFound() {
			pkg.RespondWithError(w, http.StatusNotFound, "Not Found")
			return
		}

		app.DB.Model(member).UpdateColumn(&models.Members{
			Mail:       memberQuery.Mail,
			Fullname:   memberQuery.Fullname,
			Username:   memberQuery.Username,
			Passw:      memberQuery.Passw,
			EtatCompte: "0",
			LightMode:  "on",
		})

		pkg.RespondWithNoContent(w)
	}
}

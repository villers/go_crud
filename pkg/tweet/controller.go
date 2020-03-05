package tweet

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_crud/pkg"
	"go_crud/pkg/container"
	"go_crud/pkg/tweet/models"
	"net/http"
	"time"
)

func GetTweets(app *container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweets []models.Tweets
		app.DB.Find(&tweets)

		pkg.RespondWithJSON(w, http.StatusOK, tweets)
	}
}

func GetTweet(app *container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var tweets models.Tweets

		if app.DB.First(&tweets, vars["id"]).RecordNotFound() {
			pkg.RespondWithError(w, http.StatusNotFound, "Not Found")
			return
		}

		pkg.RespondWithJSON(w, http.StatusOK, tweets)
	}
}

func PostTweet(app *container.Container) http.HandlerFunc {

	type jsonQuery struct {
		IdUser  uint   `json:"id_user"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var tweet jsonQuery
		err := json.NewDecoder(r.Body).Decode(&tweet)

		if err != nil {
			pkg.RespondWithError(w, http.StatusBadRequest, "Bad Request "+err.Error())
			return
		}

		app.DB.Create(&models.Tweets{
			IdUser:      tweet.IdUser,
			Message:     tweet.Message,
			DateSent:    time.Now(),
			FavCounter:  0,
			RtCounter:   0,
			CommCounter: 0,
		})

		pkg.RespondWithNoContent(w)
	}
}

func UpdateTweet(app *container.Container) http.HandlerFunc {

	type jsonQuery struct {
		Message     string `json:"message"`
		FavCounter  uint   `json:"fav_counter"`
		RtCounter   uint   `json:"rt_counter"`
		CommCounter uint   `json:"comm_counter"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var tweetQuery jsonQuery
		err := json.NewDecoder(r.Body).Decode(&tweetQuery)

		if err != nil {
			pkg.RespondWithError(w, http.StatusBadRequest, "Bad Request "+err.Error())
			return
		}

		vars := mux.Vars(r)

		var tweet models.Tweets
		if app.DB.First(&tweet, vars["id"]).RecordNotFound() {
			pkg.RespondWithError(w, http.StatusNotFound, "Not Found")
			return
		}

		app.DB.Model(tweet).UpdateColumn(&models.Tweets{
			Message:     tweetQuery.Message,
			DateSent:    time.Now(),
			FavCounter:  tweetQuery.FavCounter,
			RtCounter:   tweetQuery.RtCounter,
			CommCounter: tweetQuery.CommCounter,
		})

		pkg.RespondWithNoContent(w)
	}
}

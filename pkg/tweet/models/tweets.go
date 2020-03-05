package models

import (
	"time"
)

type Tweets struct {
	IdTweet     uint      `gorm:"primary_key;auto_increment:true" json:"id_tweet"`
	IdUser      uint      `json:"id_user"`
	Message     string    `json:"message"`
	DateSent    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_sent"`
	FavCounter  uint      `json:"fav_counter"`
	RtCounter   uint      `json:"rt_counter"`
	CommCounter uint      `json:"comm_counter"`
}

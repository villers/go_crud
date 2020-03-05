package models

import "time"

type Members struct {
	Id              uint   `gorm:"primary_key;auto_increment:true" json:"id"`
	Mail            string `json:"mail"`
	Fullname        string `son:"fullname"`
	Username        string `json:"username"`
	Passw           string
	DateInscription time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"data_inscription"`
	DateNaissance   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_naissance"`
	EtatCompte      string    `json:"etat_compte"`
	Avatar          string    `json:"avatar"`
	Banner          string    `json:"banner"`
	Pays            string    `json:"pays"`
	Ville           string    `json:"ville"`
	Biography       string    `json:"biography"`
	Genre           string    `json:"genre"`
	Tel             string    `json:"tel"`
	SiteWeb         string    `json:"site_web"`
	LightMode       string    `json:"light_mode"`
}

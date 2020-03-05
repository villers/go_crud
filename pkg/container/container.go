package container

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Container struct {
	Router *mux.Router
	DB     *gorm.DB
	Debug  bool
}

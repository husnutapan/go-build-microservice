package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/husnutapan/go-build-microservice/pojo"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	Driver   = os.Getenv("DRIVER")
	PUser    = os.Getenv("USER")
	Password = os.Getenv("PASS")
	Port     = os.Getenv("PORT")
	Host     = os.Getenv("HOST")
	DbName   = os.Getenv("DB")
)

type ServerInformations struct {
	Database *gorm.DB
	Router   *mux.Router
}

func (svr *ServerInformations) Connection() {
	ObtainDBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", Host, Port, PUser, DbName, Password)

	var err error
	svr.Database, err = gorm.Open(Driver, ObtainDBURL)

	if err != nil {
		fmt.Println("Cannot connect to database")
	} else {
		fmt.Println("We are connected to the database")
	}

	LoadDatabase(svr.Database)
}

func LoadDatabase(db *gorm.DB) {
	err := db.DropTableIfExists(&pojo.User{}).Error

	if err != nil {
		fmt.Println("Occur error while dropping table")
	}
	err = db.Debug().AutoMigrate(&pojo.User{}).Error

	if err != nil {
		log.Println("Cannot migrate table")
	}
}

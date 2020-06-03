package utility

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/husnutapan/go-build-microservice/handler"
	"github.com/husnutapan/go-build-microservice/pojo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type DBInfo struct {
	driver   string
	pUser    string
	password string
	port     string
	host     string
	dbName   string
}

var dbInfo *DBInfo

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
	dbInfo = &DBInfo{
		driver:   os.Getenv("DRIVER"),
		pUser:    os.Getenv("DBUSER"),
		password: os.Getenv("PASS"),
		port:     os.Getenv("PORT"),
		host:     os.Getenv("HOST"),
		dbName:   os.Getenv("DB"),
	}
}

type ServerInformations struct {
	Database *gorm.DB
	Router   *mux.Router
}

func (svr *ServerInformations) UpServer() {
	ServerRoutings(svr)
	Connection(svr)
	http.ListenAndServe(":8080", svr.Router)
}

func ServerRoutings(svr *ServerInformations) {
	svr.Router = mux.NewRouter()
	svr.Router.HandleFunc("/", AddHeaderToJSON(handler.Home)).Methods("GET")
}

//database connection
func Connection(svr *ServerInformations) {

	ObtainDBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbInfo.host, dbInfo.port, dbInfo.pUser, dbInfo.dbName, dbInfo.password)

	var err error
	svr.Database, err = gorm.Open(dbInfo.driver, ObtainDBURL)

	if err != nil {
		fmt.Println("Cannot connect to database")
	} else {
		fmt.Println("We are connected to the database")
	}

	LoadDatabase(svr.Database)
}

//create or delete table when up server through orm tool
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

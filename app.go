package parking_API

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}




func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/zones", a.getZones).Methods("GET")
	a.Router.HandleFunc("/zone", a.createZone).Methods("POST")
	a.Router.HandleFunc("/zone/{id:[0-9]+}", a.getZone).Methods("GET")
	a.Router.HandleFunc("/zone/{id:[0-9]+}", a.updateZone).Methods("PUT")
	a.Router.HandleFunc("/zone/{id:[0-9]+}", a.deleteZone).Methods("DELETE")
	////////////
	a.Router.HandleFunc("/blocks", a.getBlocks).Methods("GET")
	a.Router.HandleFunc("/block", a.createBlock).Methods("POST")
	a.Router.HandleFunc("/block/{id:[0-9]+}", a.getBlock).Methods("GET")
	a.Router.HandleFunc("/block/{id:[0-9]+}", a.updateBlock).Methods("PUT")
	a.Router.HandleFunc("/block/{id:[0-9]+}", a.deleteBlock).Methods("DELETE")
	/////////////
	a.Router.HandleFunc("/cars", a.getCars).Methods("GET")
	a.Router.HandleFunc("/car", a.createCar).Methods("POST")
	a.Router.HandleFunc("/car/{id:[0-9]+}", a.getCar).Methods("GET")
	a.Router.HandleFunc("/car/{id:[0-9]+}", a.updateCar).Methods("PUT")
	a.Router.HandleFunc("/car/{id:[0-9]+}", a.deleteCar).Methods("DELETE")

}
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "golang.org/x/text/date"
)

//Db configuration
var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"Ashish:AshishDB@tcp(easytripz.cscqq6zfyvxt.ap-south-1.rds.amazonaws.com:3306)/busbookingsystemdb")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	InitDB()
	defer db.Close()
	log.Println("Starting the HTTP server on port 9080")
	router := mux.NewRouter()

	router.HandleFunc("/buses",
		GetBuses).Methods("GET")
	router.HandleFunc("/buses/{bus_no}",
		GetBusById).Methods("GET")
	router.HandleFunc("/buses",
		AddBus).Methods("POST")
	router.HandleFunc("/buses/{bus_no}",
		UpdateBus).Methods("PUT")
	router.HandleFunc("/buses/{bus_no}",
		DeleteBus).Methods("DELETE")

	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})

}

//Get List of Buses
func GetBuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buses []Bus
	result, err := db.Query("SELECT Bus_no, Bus_name," +
		"Bus_type,Src,Dest,Arrival,Departure from Bus_details")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	for result.Next() {
		var bus Bus
		//check condition of src, dest and date entrered by user same or not
		//If same then show that bus in list else don't
		err := result.Scan(&bus.Bus_no, &bus.Bus_name,
			&bus.Bus_type, &bus.Src, &bus.Dest, &bus.Arrival, &bus.Departure)
		if err != nil {
			panic(err.Error())
		}

		if keyVal["src"] == bus.Src && keyVal["dest"] == bus.Dest {
			buses = append(buses, bus)
		}
	}

	json.NewEncoder(w).Encode(buses)
}

//Get Bus by ID
func GetBusById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT Bus_no, Bus_name,"+
		"Bus_type,Arrival,Departure from Bus_details WHERE bus_no = ?", params["bus_no"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var bus Bus
	for result.Next() {
		err := result.Scan(&bus.Bus_no, &bus.Bus_name,
			&bus.Bus_type, &bus.Arrival, &bus.Departure)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(bus)
}

//Add Bus
func AddBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO Bus_details(bus_no," +
		"bus_name,bus_type,src,dest,arrival,departure,Capacity,No_of_available_seats)" +
		"VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	bus_no := keyVal["bus_no"]
	bus_name := keyVal["bus_name"]
	bus_type := keyVal["bus_type"]
	src := keyVal["src"]
	dest := keyVal["dest"]
	Arrival := keyVal["arrival"]
	Departure := keyVal["departure"]
	Capacity := keyVal["capacity"]
	No_of_available_seats := keyVal["No_of_available_seats"]

	_, err = stmt.Exec(bus_no, bus_name, bus_type, src, dest, Arrival, Departure, Capacity, No_of_available_seats)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New Bus was Added!")
}

//Update Bus Details
func UpdateBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE Bus_details SET bus_no = ?," +
		"bus_name= ?, bus_type=?, src=?, dest=?, arrival=?, departure=?, capacity=?, no_of_available_seats=? WHERE bus_no = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	bus_no := keyVal["bus_no"]
	bus_name := keyVal["bus_name"]
	bus_type := keyVal["bus_type"]
	src := keyVal["src"]
	dest := keyVal["dest"]
	Arrival := keyVal["arrival"]
	Departure := keyVal["departure"]
	Capacity := keyVal["capacity"]
	No_of_available_seats := keyVal["No_of_available_seats"]
	_, err = stmt.Exec(bus_no, bus_name, bus_type, src, dest, Arrival, Departure, Capacity, No_of_available_seats,
		params["bus_no"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Bus with Bus_no = %s was updated",
		params["Bus_no"])
}

//Remove Bus
func DeleteBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bus_id := params["bus_no"]
	stmt, err := db.Prepare("DELETE FROM Bus_details WHERE Bus_no = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["bus_no"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Bus with bus_no = %s was deleted", bus_id)
}

type Bus struct {
	Bus_no                int    `json:"bus_no"`
	Bus_name              string `json:"bus_name"`
	Bus_type              string `json:"bus_type"`
	Src                   string `json:"src"`
	Dest                  string `json:"dest"`
	Arrival               string `json:"arrival"`
	Departure             string `json:"departure"`
	Capacity              int    `json:"capacity"`
	No_of_available_seats int    `json:"no_of_available_seats"`
}

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}

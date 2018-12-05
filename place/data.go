package place

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fannyhasbi/bus-schedule-rest-go/data"
)

func ReturnPlaces(w http.ResponseWriter, r *http.Request) {
	var place Place
	var arr_places []Place
	var response ResponsePlace

	db := data.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tempat")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&place.Id, &place.Nama); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_places = append(arr_places, place)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_places

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

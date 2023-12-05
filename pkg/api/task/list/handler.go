package function

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	database "github.com/tommoulard/side/pkg/db"
	"github.com/tommoulard/side/pkg/openfaassdk"
)

var db *database.DB

func init() {
	mongoUsername, err := openfaassdk.ReadSecret("mongo-db-username")
	if err != nil {
		panic(err)
	}
	mongoPassword, err := openfaassdk.ReadSecret("mongo-db-password")
	if err != nil {
		panic(err)
	}

	db, err = database.New(os.Getenv("MONGO_HOST"), mongoUsername, mongoPassword)
	if err != nil {
		panic(err)
	}
}

type Task struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Organisation struct {
		Name       string `json:"name"`
		Address    string `json:"address"`
		PictureURL string `json:"picture_url"`
	} `json:"organisation"`
	Shifts []struct {
		ID        string `json:"id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		Slots     struct {
			Filled int `json:"filled"`
			Total  int `json:"total"`
		} `json:"slots"`
	} `json:"shifts"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	if err := db.Find("tasks", nil, &tasks); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("failed to find tasks: %v", err)

		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("failed to encode tasks: %v", err)

		return
	}
}

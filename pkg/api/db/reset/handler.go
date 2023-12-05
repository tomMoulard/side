package function

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	database "github.com/tommoulard/side/pkg/db"
	"github.com/tommoulard/side/pkg/openfaassdk"
)

var (
	//go:embed fixtures
	fixtures embed.FS

	db *database.DB
)

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

func Handle(w http.ResponseWriter, r *http.Request) {
	dir, _ := fixtures.ReadDir("fixtures")
	for _, f := range dir {
		fmt.Println(f.Name())
		file, err := fixtures.Open("fixtures/" + f.Name())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to open file %q: %v", f.Name(), err)

			return
		}

		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to read file %q: %v", f.Name(), err)

			return
		}

		var v []interface{}
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to unmarshal data: %v", err)

			return
		}

		collection := strings.TrimSuffix(f.Name(), ".json")
		if err := db.InsertMany(collection, v); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to insert data: %v", err)

			return
		}
	}
}

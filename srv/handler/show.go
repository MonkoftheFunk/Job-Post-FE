package handler

import (
	repo "Job-Post-FE/srv/mongo"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type listing struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

func HandleShow(i *inertia.Inertia, config *repo.Config, l *logrus.Logger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		repo := repo.NewClient(config)
		coll := repo.Database(config.Database).Collection("listings")
		var result listing
		err := coll.FindOne(context.TODO(), bson.D{{"slug", vars["slug"]}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the slug %s\n", vars["slug"])
			return
		}
		if err != nil {
			panic(err)
		}

		err = i.Render(w, r, "Listing/Show", inertia.Props{
			"listing": result,
		})
		if err != nil {
			panic(err)
		}
	}

	return http.HandlerFunc(fn)
}

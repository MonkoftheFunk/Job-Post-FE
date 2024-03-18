package handler

import (
	repo "Job-Post-FE/srv/mongo"
	"context"
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func HandleIndex(i *inertia.Inertia, config *repo.Config, l *logrus.Logger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		repo := repo.NewClient(config)
		coll := repo.Database(config.Database).Collection("listings")
		var results []listing
		var query bson.M
		search := r.URL.Query().Get("search")
		if search != "" {
			like := "/" + search + "/i"
			query = bson.M{"$or": []bson.M{{"title": like}, {"company": like}, {"location": like}}}
		}
		activeTag := r.URL.Query().Get("activeTag")
		if activeTag != "" {
			query["tags"] = []string{activeTag}
		}

		// todo log and handle
		cursor, err := coll.Find(context.Background(), query)
		if err != nil {
			l.Fatal(err)
		}

		err = cursor.All(context.Background(), &results)
		if err != nil {
			l.Fatal(err)
		}
		if results == nil {
			results = []listing{}
		}

		err = i.Render(w, r, "Listing/Index", inertia.Props{
			"listings":   results,
			"tags":       []string{}, // todo query all stored tags
			"active_tag": activeTag,
			"search":     search,
		})
		if err != nil {
			l.Fatal(err)
		}
	})
}

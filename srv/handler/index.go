package handler

import (
	repo "Job-Post-FE/srv/mongo"
	"Job-Post-FE/srv/session"
	"context"
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func HandleIndex(i *inertia.Inertia, mconfig *repo.Config, sconfig *session.Config, l *logrus.Logger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		session, err := session.NewClient(sconfig).Get(r)
		if err != nil {
			l.Fatal(err)
		}
		//todo
		if session["_token"] != nil {
			l.Infof(session["_token"].(string))
		}

		repo := repo.NewClient(mconfig)
		coll := repo.Database(mconfig.Database).Collection("listings")
		var results []listing
		var query bson.M
		search := r.URL.Query().Get("search")
		if search != "" {
			regexPattern := primitive.Regex{Pattern: search, Options: "i"}
			query = bson.M{"$or": []bson.M{
				{"title": regexPattern},
				{"company": regexPattern},
				{"location": regexPattern},
			}}
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

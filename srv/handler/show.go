package handler

import (
	repo "Job-Post-FE/srv/mongo"
	"Job-Post-FE/srv/session"
	"context"
	"github.com/gorilla/mux"
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type listing struct {
	Slug          string `json:"slug"`
	Title         string `json:"title"`
	Company       string `json:"company"`
	Content       string `json:"content"`
	IsActive      bool   `json:"is_active"`
	IsHighlighted bool   `json:"is_highlighted"`
	Link          string `json:"link"`
	Location      string `json:"location"`
	ClicksCount   string `json:"clicksCount"`
	LogoUri       string `json:"logoUri"`
	SinceCreated  string `json:"sinceCreated"`
	TagsCSV       string `json:"tagsCSV"`
}

func HandleShow(i *inertia.Inertia, mconfig *repo.Config, sconfig *session.Config, l *logrus.Logger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		session, err := session.NewClient(sconfig).Get(r)
		if err != nil {
			l.Fatal(err)
		}
		//todo
		l.Infof(session["_token"].(string))

		repo := repo.NewClient(mconfig)
		coll := repo.Database(mconfig.Database).Collection("listings")
		var result listing

		// todo log and handle
		err = coll.FindOne(context.TODO(), bson.D{{"slug", vars["slug"]}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			l.Warning("No document was found with the slug %s\n", vars["slug"])
			http.NotFound(w, r)
			return
		}
		if err != nil {
			l.Fatal(err)
		}

		err = i.Render(w, r, "Listing/Show", inertia.Props{
			"listing": result,
		})
		if err != nil {
			l.Fatal(err)
		}
	})
}

package handler

import (
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"net/http"
)

type listing struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

func HandleShow(i *inertia.Inertia, l *logrus.Logger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)
		//vars["slug"]
		_ = i.Render(w, r, "Listing/Show", inertia.Props{
			"listing": listing{
				Slug:  "test",
				Title: "Test Listing",
			},
		})
		/*if err != nil {
			handleServerErr(w, err)
		}*/
	}

	return http.HandlerFunc(fn)
}

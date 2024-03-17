package handler

import (
	"Job-Post-FE/srv/mongo"
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleIndex(i *inertia.Inertia, config *mongo.Config, l *logrus.Logger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		_ = i.Render(w, r, "Listing/Index", inertia.Props{
			"search":     "",
			"active_tag": "",
			"tags":       []string{"d", "s"},
			"listings":   []listing{},
		})

		/*if err != nil {
			handleServerErr(w, err)
		}*/
	}

	return http.HandlerFunc(fn)
}

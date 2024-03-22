package server

import (
	"Job-Post-FE/srv/handler"
	"Job-Post-FE/srv/mongo"
	"Job-Post-FE/srv/session"
	"github.com/gorilla/mux"
	inertia "github.com/romsar/gonertia"
	"github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	port          int
	logger        *logrus.Logger
	router        *mux.Router
	inertia       *inertia.Inertia
	assets        fs.FS
	mongoConfig   *mongo.Config
	sessionConfig *session.Config
}

func Run(port int, fs fs.FS, mconfig *mongo.Config, sconfig *session.Config) error {
	i, err := inertia.New("./resources/app.html")
	if err != nil {
		return err
	}

	s := server{
		logger:        logrus.New(),
		port:          port,
		router:        mux.NewRouter(),
		inertia:       i,
		assets:        fs,
		mongoConfig:   mconfig,
		sessionConfig: sconfig,
	}

	s.bindRoutes()
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(s.port),
		Handler:      s.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *server) bindRoutes() {
	// todo use 1.22 mux routing
	// todo refactor for better testing and mocks
	s.router.PathPrefix("/dist/").Handler(http.FileServer(http.FS(s.assets)))
	s.router.Handle("/l/{slug}", s.inertia.Middleware(handler.HandleShow(s.inertia, s.mongoConfig, s.sessionConfig, s.logger)))
	s.router.Handle("/", s.inertia.Middleware(handler.HandleIndex(s.inertia, s.mongoConfig, s.sessionConfig, s.logger)))
}

package apiDir

import (
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type Rest struct {
	address  string
	mux      *mux.Router
	listener net.Listener
}

func New(address string) *Rest {
	api := mux.NewRouter()

	api.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./statics"))))

	return &Rest{
		mux:     api,
		address: address,
	}
}

func (rest *Rest) Listen() (err error) {
	rest.listener, err = net.Listen("tcp", rest.address)
	if err != nil {
		return err
	}

	r := http.NewServeMux()
	r.Handle("/", rest.mux)

	server := &http.Server{Handler: r}

	rest.setupMiddleware()

	return server.Serve(rest.listener)
}

func (rest *Rest) setupMiddleware() {
	rest.mux.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("new request", r.RequestURI)
			handler.ServeHTTP(w, r)
		})
	})
}

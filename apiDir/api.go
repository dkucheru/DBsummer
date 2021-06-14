package apiDir

import (
	"DBsummer/serviceDir"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type Rest struct {
	address  string
	mux      *mux.Router
	listener net.Listener
	service  *serviceDir.Service
}

func New(address string, service *serviceDir.Service) *Rest {
	rest := &Rest{
		address: address,
		service: service,
	}

	api := mux.NewRouter()

	api.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("C:\\Users\\Алина\\go\\src\\DBsummer\\runDir\\staticsDir"))))
	api.HandleFunc("/test", rest.test)
	api.HandleFunc("/subjects/{id}", rest.getSubject).Methods("GET")
	api.HandleFunc("/subjects", rest.getSubjects).Methods("GET")
	api.HandleFunc("/students", rest.getStudentsList).Methods("GET")
	api.HandleFunc("/students/who_skipped_exam", rest.getAllBorjniki).Methods("GET")
	api.HandleFunc("/subjects/{fn}/{ln}/{mn}/{ye}", rest.getStudentByPIB).Methods("GET")
	api.HandleFunc("/sheets/{fn}/{ln}/{mn}/{subj}/{gr}/{ye}", rest.getSheetFromParams).Methods("GET")
	api.HandleFunc("/sheets/add", rest.postSheet).Methods("GET")
	api.HandleFunc("/subjects/by_year/{year}", rest.getSubjectsByYear).Methods("GET")
	api.HandleFunc("/teachers/{pass}", rest.getTeacherPasses).Methods("GET")

	rest.mux = api

	return rest
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

type Response struct {
	Status int
	Data   interface{}
}

func (rest *Rest) sendError(w http.ResponseWriter, err error) {
	bytes, err := json.Marshal(Response{
		Status: 2,
		Data:   err,
	})
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}

func (rest *Rest) sendData(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(Response{
		Status: 1,
		Data:   data,
	})
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}

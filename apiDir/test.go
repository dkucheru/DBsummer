package apiDir

import (
	"log"
	"net/http"
)

func (rest *Rest) test(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Daria!"))
	if err != nil {
		log.Println(err)
	}
}

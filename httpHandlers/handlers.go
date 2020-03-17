package httpHandlers

import "net/http"
import "log"
import "github.com/vvidovic/go-ws-example/httpHandlers/httpUtils"

func HandleRequestCollection(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming Collection Request, %s %v", r.Method, r.URL)
	switch r.Method {
	case http.MethodGet:
		List(w, r)
	case http.MethodPost:
		Add(w, r)
	default:
		httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
	}
}

func HandleRequestEntity(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming Entity Request, %s %v", r.Method, r.URL)
	switch r.Method {
	case http.MethodGet:
		Get(w, r)
	case http.MethodDelete:
		Remove(w, r)
	default:
		httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
	}
}

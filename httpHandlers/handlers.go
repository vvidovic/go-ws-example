package httpHandlers

import "net/http"
import "log"
import "github.com/vvidovic/go-ws-example/httpHandlers/httpUtils"

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming Request, %s %v", r.Method, r.URL)
	switch r.Method {
	case http.MethodGet:
		List(w, r)
		break
	case http.MethodPost:
		Add(w, r)
		break
	case http.MethodDelete:
		Remove(w, r)
		break
	default:
		httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

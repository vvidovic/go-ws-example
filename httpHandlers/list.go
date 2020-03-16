package httpHandlers

import (
	"net/http"

	"github.com/vvidovic/services/httpHandlers/httpUtils"
	"github.com/vvidovic/services/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.Get())
}

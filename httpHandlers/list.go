package httpHandlers

import (
	"net/http"

	"github.com/vvidovic/go-ws-example/httpHandlers/httpUtils"
	"github.com/vvidovic/go-ws-example/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.Get())
}

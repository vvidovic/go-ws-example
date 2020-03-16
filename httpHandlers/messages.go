package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"log"

	"github.com/vvidovic/go-ws-example/httpHandlers/httpUtils"
	"github.com/vvidovic/go-ws-example/storage"
	"github.com/vvidovic/go-ws-example/structs"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.List())
}

func Get(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/messages/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		httpUtils.HandleError(&w, 400, "Bad Request", "Invalid ID", nil)
		return
	}
	message := storage.Get(id)
	if message.ID == 0 {
		httpUtils.HandleError(&w, 404, "Not Found", "Non-existing ID", nil)
		return
	}
	httpUtils.HandleSuccess(&w, message)
}

func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var message structs.Message

	err = json.Unmarshal(byteData, &message)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if message.Message == "" || message.Sender == "" {
		httpUtils.HandleError(&w, 400, "Bad Request", "Unmarshalled JSON didn't have required fields", nil)
		return
	}

	id := storage.Add(message)

	log.Println("Added message:", message)

	httpUtils.HandleSuccess(&w, structs.ID{ID: id})
}

func Remove(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/messages/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		httpUtils.HandleError(&w, 400, "Bad Request", "Invalid ID", nil)
		return
	}

	if storage.Remove(id) {
		httpUtils.HandleSuccess(&w, structs.ID{ID: id})
	} else {
		httpUtils.HandleError(&w, 400, "Bad Request", "ID not found", nil)
	}
}

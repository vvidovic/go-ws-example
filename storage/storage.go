package storage

import (
	"github.com/vvidovic/go-ws-example/structs"
)

var store structs.MessageList
var currentMaxId = 1

func List() structs.MessageList {
	return store
}

func Get(id int) structs.Message {
	for _, message := range store {
		if message.ID == id {
			return message
		}
	}
	return structs.Message{}
}

func Add(message structs.Message) int {
	message.ID = currentMaxId
	currentMaxId++
	store = append(store, message)
	return message.ID
}

func Remove(id int) bool {
	index := -1

	for i, message := range store {
		if message.ID == id {
			index = i
		}
	}

	if index != -1 {
		store = append(store[:index], store[index+1:]...)
	}

	// Returns true if item was found & removed
	return index != -1
}

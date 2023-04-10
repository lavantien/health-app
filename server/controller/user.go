package controller

import (
	"encoding/json"
	"health-app/server/common"
	"health-app/server/data"
	"health-app/server/model"
	"io/ioutil"
	"net/http"
	"sync"
)

type users []model.User

type UserHandler struct {
	sync.Mutex
	users users
}

func (uh *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		uh.get(w, r)
	case "POST":
		uh.post(w, r)
	default:
		common.ResponseWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (uh *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	defer uh.Unlock()
	uh.Lock()
	id, err := common.IDFromURL(r)
	if err != nil {
		common.ResponseWithJSON(w, http.StatusOK, uh.users)
		return
	}
	if !uh.idExists(id) {
		common.ResponseWithError(w, http.StatusNotFound, "not found")
		return
	}
	common.ResponseWithJSON(w, http.StatusOK, uh.users[id])
}

func (uh *UserHandler) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		common.ResponseWithError(w, http.StatusUnsupportedMediaType, "content-type 'application/json' required")
		return
	}
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		common.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if uh.nameExists(user.Name) {
		common.ResponseWithError(w, http.StatusConflict, "name already exists")
		return
	}
	defer uh.Unlock()
	uh.Lock()
	user.ID = len(uh.users)
	uh.users = append(uh.users, user)
	common.ResponseWithJSON(w, http.StatusOK, user)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		users: data.InitUsers,
	}
}

func (uh *UserHandler) idExists(id int) bool {
	for _, user := range uh.users {
		if user.ID == id {
			return true
		}
	}
	return false
}

func (uh *UserHandler) nameExists(name string) bool {
	for _, user := range uh.users {
		if user.Name == name {
			return true
		}
	}
	return false
}

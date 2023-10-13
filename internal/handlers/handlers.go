package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/grbarra/rtask/internal/models"
	"gitlab.com/grbarra/rtask/internal/service"
)

type Handler struct {
	service service.UserService
}

func NewHandler(service service.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	us, err := h.service.GetUser(user.Key)

	if err = json.NewEncoder(w).Encode(us); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HandlerSetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.service.SetUser(user.Key, user.Name, user.TTL)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandlerAddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.service.AddUser(user.Key, user.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandlerDelUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.service.DelUser(user.Key)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandlerKeysUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	val, err := h.service.KeysUser(user.Key)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = json.NewEncoder(w).Encode(val); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

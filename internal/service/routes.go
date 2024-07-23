package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/kafka"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/model"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/utils"
)

type Handler struct {//Handler shouldn't know exactly how message operations are implemented
	store model.MessagesStore
}

func NewHandler(store model.MessagesStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/message", h.handleMessage).Methods("POST")
	router.HandleFunc("/statistics", h.handleStatistics).Methods("GET")
}

func (h *Handler) handleMessage(w http.ResponseWriter, r *http.Request) {
	var payload model.MessagesPayload //Payload is data that we receive/send
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	id, err := h.store.CreateMessage(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//Convert payload to bytes
	messageBytes, err := json.Marshal(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	} 

	//Send bytes to kafka
	err = kafka.PushMessageToQueue("messages", messageBytes)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	h.store.MarkMessageAsProcessed(id)

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleStatistics(w http.ResponseWriter, r *http.Request) {
	statistics, err := h.store.GetStatistics()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, statistics)
}
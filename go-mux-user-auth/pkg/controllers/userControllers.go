package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KannanPalani57/go-mux-user-auth/pkg/models"
	"github.com/KannanPalani57/go-mux-user-auth/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}

	utils.ParseBody(r, CreateUser)

	u := CreateUser.CreateUser()

	res, _ := json.Marshal(u)

	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

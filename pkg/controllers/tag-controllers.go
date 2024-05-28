package controllers

import (
	"encoding/json"
	"github/fokaz-c/go-book-manager/pkg/models"
	"github/fokaz-c/go-book-manager/pkg/utils"
	"net/http"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	CreateTag := &models.Tag{}
	utils.ParseBody(r, CreateTag)
	b, _ := CreateTag.CreateTag()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

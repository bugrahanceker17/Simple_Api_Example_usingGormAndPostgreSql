package handlers

import (
	"GormExample/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllBook(res http.ResponseWriter, req *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(books)
}

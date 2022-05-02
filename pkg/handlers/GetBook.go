package handlers

import (
	"GormExample/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) GetBook(res http.ResponseWriter, req *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	// find book by id
	var book models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(book)
}

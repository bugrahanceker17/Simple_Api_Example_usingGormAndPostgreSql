package handlers

import (
	"GormExample/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) DeleteBook(res http.ResponseWriter, req *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	// Find the book by id
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	//Delete that book
	h.DB.Delete(&book, id)

	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("Deleted!")
}

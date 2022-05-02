package handlers

import (
	"GormExample/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(res http.ResponseWriter, req *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	defer req.Body.Close()
	// Read request body
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var update models.Book
	json.Unmarshal(body, &update)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Desc = update.Desc
	book.Title = update.Title
	book.Author = update.Author

	// Update
	h.DB.Save(&book)

	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("Updated!")
}

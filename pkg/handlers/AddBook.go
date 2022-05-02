package handlers

import (
	"GormExample/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddBook(res http.ResponseWriter, req *http.Request) {
	// Read to request body
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Book dummy
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send 201 status code
	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("Created")
}

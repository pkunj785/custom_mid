package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

)

type Lesson struct {
	Title string 
	Summery string
}

func Handle(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {

		var tempLesson Lesson

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempLesson)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()

		fmt.Printf("Titel: %s. Summery: %s\n ", tempLesson.Title, tempLesson.Summery)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not Allowed"))
	}	
}
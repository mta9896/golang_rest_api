package httprest

import (
	"encoding/json"
	"fmt"
	"io"
	"mta9896/restapi/internal/database"
	"mta9896/restapi/internal/entity"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/items", getItemsHandler).Methods("GET")
	router.HandleFunc("/items", createItemHandler).Methods("POST")

	port := 8080
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, _ := database.FetchAllItems()


	response, err := json.Marshal(items)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ , err = w.Write(response)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var newItem entity.Item
	err = json.Unmarshal(body, &newItem)
	if err != nil {
		http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}

	err = database.InsertItem(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
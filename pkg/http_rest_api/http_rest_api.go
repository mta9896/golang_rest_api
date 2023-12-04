package httprestapi

import(
	"encoding/json"
	"net/http"
	"io"
	"fmt"
	"github.com/gorilla/mux"
	"mta9896/golang_rest_api/pkg/crud"
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
	response, err := json.Marshal(crud.List())

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var newItem crud.Item
	err = json.Unmarshal(body, &newItem)
	if err != nil {
		http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}

	crud.Create(newItem)

	response, err := json.Marshal(newItem)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
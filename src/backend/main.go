package main

import (
	"encoding/json"
	back "app/lib"
	"fmt"
	"io/ioutil"
	"net/http"
)

var gotInfo back.LinkInfo

var res []string

var exec int64

func handleInsert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(requestBody, &gotInfo); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Request body:", string(requestBody))

	// fmt.Printf("Received data: %+v\n", pointDataGlobale)

	// var time_string string;
	res,exec = back.Back_Main(gotInfo)
	
	tempResponse := struct {
		Lencs []string `json:"links"`
		Exec int64 `json:"exec"`
	}{
		Lencs: res,
		Exec: exec,
	}
	// responseData := struct {
	// 	Message     string   `json:"message"`
	// 	GottenPoints []backend.Point `json:"gottenPoint"`
	// 	ControlPoints []backend.Point `json:"controlPoint"`
	// 	Time_String string `json:"time_string"`
	// }{
	// 	Message:     "Received data successfully",
	// 	GottenPoints: gottenPoint,
	// 	ControlPoints: controlPoint,
	// 	Time_String: time_string,
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempResponse)
}

func main() {
	http.HandleFunc("/", handleInsert)
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}

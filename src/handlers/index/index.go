package index

import (
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := &IndexResponse{
		Status: 200,
		Message: "Welcome to v1 of the sugoiart API",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
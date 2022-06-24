package art

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := &IndexResponse{
		Status: 200,
		Message: "This is /api/art",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
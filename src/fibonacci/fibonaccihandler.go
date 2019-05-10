package fibonacci

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type response struct {
	N      int    `json:"n"`
	Result string `json:"result"`
}

func (r response) json() ([]byte, error) {
	return json.Marshal(
		response{
			r.N,
			r.Result,
		},
	)
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.Header().Add("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	n, err := strconv.Atoi(req.URL.Query().Get("fib"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := ComputeNthFibonacci(n)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := response{n, result.String()}.json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	_, _ = w.Write(response)
}

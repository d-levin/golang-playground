package httpserver

import (
	"encoding/json"
	"errors"
	"fibonacci"
	"net/http"
	"strconv"
	"strings"
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

func fibonacciHandler(w http.ResponseWriter, req *http.Request) {
	n, err := extractN(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := fibonacci.ComputeNthFibonacci(n)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	response, err := response{n, result.String()}.json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}

func extractN(req *http.Request) (int, error) {
	parts := strings.Split(req.URL.RawQuery, "=")

	if len(parts) < 2 || parts[0] != "fib" {
		return -1, errors.New("invalid param")
	}

	n, err := strconv.Atoi(parts[1])

	if err != nil {
		return -1, errors.New("invalid value")
	}

	return n, nil
}

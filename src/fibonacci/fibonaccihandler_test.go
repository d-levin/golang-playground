package fibonacci

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testhelpers"
	"testing"
)

func Test_Handler_GivenValidRequest_ShouldReturnNthFibonacci(t *testing.T) {
	pairs := []testhelpers.Pair{
		{0, "0"},
		{1, "1"},
		{2, "1"},
		{3, "2"},
		{4, "3"},
		{5, "5"},
		{6, "8"},
		{7, "13"},
		{8, "21"},
		{9, "34"},
		{10, "55"},
		{100, "354224848179261915075"},
		{200, "280571172992510140037611932413038677189525"},
		{300, "222232244629420445529739893461909967206666939096499764990979600"},
	}

	for _, p := range pairs {
		req, err := http.NewRequest(http.MethodGet, "?fib="+strconv.Itoa(p.N), nil)

		if err != nil {
			t.Error(err)
		}

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(Handler)

		handler.ServeHTTP(responseRecorder, req)

		actual := response{}
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &actual)
		if actual.N != p.N {
			t.Error("Expected "+strconv.Itoa(p.N)+", got", actual.N)
		}
		if actual.Result != p.Expected {
			t.Error("Expected "+p.Expected+", got", actual.Result)
		}
	}
}

func Test_Handler_GivenValidRequest_ShouldReturnHttpStatus200(t *testing.T) {
	urls := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		100,
		200,
		300,
	}

	for _, n := range urls {
		req, err := http.NewRequest(http.MethodGet, "?fib="+strconv.Itoa(n), nil)

		if err != nil {
			t.Error(err)
		}

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(Handler)

		handler.ServeHTTP(responseRecorder, req)

		status := responseRecorder.Code
		if status != http.StatusOK {
			t.Error("Expected "+strconv.Itoa(http.StatusOK)+", got", status)
		}
	}
}

func Test_Handler_GivenValidRequest_ShouldReturnContentTypeApplicationJson(t *testing.T) {
	urls := []int{
		0,
	}

	for _, n := range urls {
		req, err := http.NewRequest(http.MethodGet, "?fib="+strconv.Itoa(n), nil)

		if err != nil {
			t.Error(err)
		}

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(Handler)

		handler.ServeHTTP(responseRecorder, req)

		header := responseRecorder.Header().Get("content-type")
		if header == "" {
			t.Error("Expected application/json, got", header)
		}
	}
}

func Test_Handler_GivenInvalidRequest_ShouldReturnHttpStatus400(t *testing.T) {
	urls := []string{
		"",
		"/",
		"/fib",
		"/fib=2",
		"?",
		"?fib",
		"?fib=",
		"?fib=a",
		"?fib=-1",
		"?a",
		"?a=",
		"?a=5",
	}

	for _, url := range urls {
		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			t.Error(err)
		}

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(Handler)

		handler.ServeHTTP(responseRecorder, req)

		status := responseRecorder.Code
		if status != http.StatusBadRequest {
			t.Error("Expected "+strconv.Itoa(http.StatusBadRequest)+", got", status)
		}
	}
}

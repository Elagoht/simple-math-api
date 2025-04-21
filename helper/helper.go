package helper

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetQueryParamNumbers(r *http.Request, keys ...string) ([]float64, error) {
	queryParams := r.URL.Query()
	numbers := make([]float64, len(keys))
	for i, key := range keys {
		value := queryParams.Get(key)
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		numbers[i] = num
	}
	return numbers, nil
}

func HandleMathController(path string, operation func(a, b float64) (float64, error)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		numbers, err := GetQueryParamNumbers(r, "a", "b")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := operation(numbers[0], numbers[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(fmt.Appendf(nil, "%f", result))
	})
}

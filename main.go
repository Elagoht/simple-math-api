package main

import (
	"fmt"
	"net/http"

	"github.com/Elagoht/simple-math-api/helper"
	"github.com/Elagoht/simple-math-api/operations"
)

func main() {
	fmt.Println("Server starting on port 8080...")
	helper.HandleMathController("/add", operations.Add)
	helper.HandleMathController("/subtract", operations.Subtract)
	helper.HandleMathController("/multiply", operations.Multiply)
	helper.HandleMathController("/divide", operations.Divide)
	helper.HandleMathController("/power", operations.Power)
	fmt.Println("Server started on port 8080")

	http.ListenAndServe(":8080", nil)
}

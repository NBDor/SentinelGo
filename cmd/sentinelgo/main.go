package main

import "fmt"

func main() {
	var int1 int = 10
	var int2 = 20
	int3 := 30

	fmt.Println(int1)
	fmt.Println(int2)
	fmt.Println(int3)

	var numerator int = 10
	var denominator int = 3

	var result, remiander int = intDivision(numerator, denominator)
	fmt.Printf("The result of the division is %v with a remainder of %v\n", result, remiander)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}

// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func main() {
//     log.Println("Starting SentinelGO service...")

//     http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "SentinelGO is healthy!")
//     })

//     log.Println("Server starting on :8080")
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         log.Fatalf("Server failed to start: %v", err)
//     }
// }

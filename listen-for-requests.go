package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting web server on port 8080 ")
	http.ListenAndServe(":8000", nil)
}

package main
import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the GO language!")
}

func main() {
		http.HandleFunc("/", handler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

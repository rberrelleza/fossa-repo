package cmd

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting hello-world server on port 8080...")
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received")
	fmt.Fprint(w, "Hello fig!")
}

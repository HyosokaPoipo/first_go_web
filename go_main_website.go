package main

import(
    "fmt"
    "log"
    "net/http"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hellow Maryadi </h1>")
}

func main() {
	http.HandleFunc("/", baseHandler)
	log.Fatal(http.ListenAndServe(":1111", nil))
}
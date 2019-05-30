package first_webapp

import (
	"fmt"
	"net/http"
)

func handler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "Hello World, %s!", request.URL.Path[1])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8088", nil)
}

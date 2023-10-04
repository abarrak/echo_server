package main

import (
	"fmt"
	"io"
	"net/http"
)

// Echo Server sample implementation.
func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<b> Hello World! </b>")
}

func headersEchoer(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%+s : %+s\n", name, h)
		}
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "\n %+v", string(body))
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/echo", headersEchoer)
	http.ListenAndServe(":8080", nil)
}

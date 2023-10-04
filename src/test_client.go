package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	ECHO_SERVER = "http://localhost:8080/echo"
)

func main() {
	// GET
	response, err := http.Get(ECHO_SERVER)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response HTTP Status: ", response.Status)
	scanner := bufio.NewScanner(response.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// POST
	values := map[string]string{"name": "John", "age": "31"}
	json_data, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	req := bytes.NewBuffer(json_data)
	resp, _ := http.Post(ECHO_SERVER, "application/json", req)
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		panic(err)
	}
}

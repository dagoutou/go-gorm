package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := new(http.Client)
	request, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	n, err := client.Do(request)

	if err != nil {
		return
	}
	body := n.Body
	all, err := io.ReadAll(body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
}

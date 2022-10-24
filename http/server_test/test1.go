package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)
}
func test(resp http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resp.Write([]byte("我收到你的请求了给你返回GET！"))
		return
	case "POST":
		a, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		head := resp.Header()
		head.Set("aaa", "7777")
		head["bbbb"] = []string{"gouzi,gouzi"}
		resp.WriteHeader(http.StatusOK)
		resp.Write(a)
		return

	}
}

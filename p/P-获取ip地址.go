package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", ExampleHandler)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": getIP(r),
	})
	w.Write(resp)

}

func getIP(r *http.Request) string {
	forwareded := r.Header.Get("X-FORWARDED-FOR")
	if forwareded != "" {
		return forwareded
	}
	return r.RemoteAddr
}

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

var urlStore = struct {
	sync.RWMutex
	mapping map[string]string
}{mapping: make(map[string]string)}

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

func generateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyz123456"
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]

	}
	return string(code)

}
func shortHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		URL string `json:"url"`
	}
	type Response struct {
		ShortURL string `json:"short_url"`
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || !strings.HasPrefix(req.URL, "http") {
		http.Error(w, "Invaild request", http.StatusBadRequest)
		return
	}
	code := generateShortCode()
	urlStore.Lock()
	urlStore.mapping[code] = req.URL
	urlStore.Unlock()

	resp := Response{ShortURL: "http://localhost:8080/" + code}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	urlStore.RLock()
	originalURL, found := urlStore.mapping[code]
	urlStore.RUnlock()

	if found {
		http.Redirect(w, r, originalURL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}

}
func main() {
	http.HandleFunc("/shorten", shortHandler)
	http.HandleFunc("/", redirectHandler)

	log.Println("Server started at :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

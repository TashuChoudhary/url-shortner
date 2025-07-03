package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

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
	//urlStore.Lock()
	//urlStore.mapping[code] = req.URL
	//urlStore.Unlock()
	if err := saveURL(code, req.URL); err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	resp := Response{ShortURL: baseURL + "/" + code}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	//urlStore.RLock()
	//originalURL, found := urlStore.mapping[code]
	//urlStore.RUnlock()
	originalURL, found := getOriginalURL(code)

	if found {
		http.Redirect(w, r, originalURL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "Welcome to URL Shortner Service 🎉 POST /shorten with JSON ")
		return

	}
	redirectHandler(w, r)
}

func main() {
	initDB()
	http.HandleFunc("/shorten", shortHandler)
	http.HandleFunc("/", rootHandler)

	log.Println("Server started at :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	sourceURL  = "https://updates.jenkins.io/update-center.json"
	targetText = "https://updates.jenkins.io/download/"
	// Change to your own mirror site
	replacementText = "https://ftp.yz.yamagata-u.ac.jp/pub/misc/jenkins/"
)

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	// 抓取原始 JSON 資料
	resp, err := http.Get(sourceURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "Failed to close response body", http.StatusInternalServerError)
			return
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Failed to read data", http.StatusInternalServerError)
		return
	}

	// 進行字串取代
	modifiedData := strings.ReplaceAll(string(data), targetText, replacementText)

	if err != nil {
		http.Error(w, "Failed to replace text", http.StatusInternalServerError)
		return
	}

	// 回傳修改後的 JSON 資料
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(modifiedData))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/update-center.json", handleAPIRequest)
	fmt.Println("Server is running on port 8080")
	fmt.Println("Browse http://YOUR_IP:8080/update-center.json to see the result.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// auth_helper/main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	metadataURLBase = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity"
)

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	audience := r.URL.Query().Get("audience")
	if audience == "" {
		log.Println("Error: Missing 'audience' query parameter")
		http.Error(w, "Missing audience query parameter", http.StatusBadRequest)
		return
	}

	fullMetadataURL := fmt.Sprintf("%s?audience=%s", metadataURLBase, audience)
	log.Printf("Fetching token for audience: %s from %s\n", audience, fullMetadataURL)

	client := &http.Client{
		Timeout: 5 * time.Second, // Set a timeout for the request
	}
	req, err := http.NewRequest("GET", fullMetadataURL, nil)
	if err != nil {
		log.Printf("Error creating request for metadata server: %v\n", err)
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	req.Header.Add("Metadata-Flavor", "Google")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error requesting token from metadata server for audience %s: %v\n", audience, err)
		http.Error(w, fmt.Sprintf("Error fetching token: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Error from metadata server (audience: %s): status %d, body: %s\n", audience, resp.StatusCode, string(bodyBytes))
		http.Error(w, fmt.Sprintf("Metadata server error: %s", resp.Status), resp.StatusCode)
		return
	}

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading token response body (audience: %s): %v\n", audience, err)
		http.Error(w, "Error reading token response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Auth-Token", string(token))
	w.Header().Set("Content-Type", "application/json") // Nginx auth_request doesn't care much about body for 2xx
	fmt.Fprintln(w, `{"success":true, "message":"Token ready to be captured from header"}`)
	log.Printf("Successfully fetched and set token for audience: %s\n", audience)
}

func main() {
	helperPort := os.Getenv("HELPER_PORT")
	if helperPort == "" {
		helperPort = "8081" // Default port
	}

	http.HandleFunc("/get-token", getTokenHandler)
	log.Printf("Starting Go auth helper service on port %s\n", helperPort)
	if err := http.ListenAndServe(":"+helperPort, nil); err != nil {
		log.Fatalf("Failed to start Go auth helper server: %v", err)
	}
}

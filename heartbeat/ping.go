package heartbeat

import (
	"log"

	"net/http"
	"net/url"
)

//TODO: Instead of sending the url of type URL we can send the string of the
//healthCheck endpoint - it would streamline the test and keep logic outside the
//function.
func PingServer(url *url.URL) bool {

	log.Printf("Requesting /healhCheck from: %s\n", url.Host)

	//get request to server /healthCheck path
	resp, err := http.Get("http://" + url.Host + "/healthCheck")
	if err != nil {
		log.Println("Server is offline, error:", err)
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

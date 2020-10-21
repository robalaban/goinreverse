package heartbeat

import (
	"github.com/tcnksm/go-httpstat"
	"log"
	"net/http"
	"net/url"
	"time"
)

//TODO: Instead of sending the url of type URL we can send the string of the
//healthCheck endpoint - it would streamline the test and keep logic outside the
//function.
func PingServer(url *url.URL) (bool, int) {
	log.Printf("Requesting /healhCheck from: %s\n", url.Host)

	req, _ := http.NewRequest("GET", "http://" + url.Host + "/healthCheck", nil)
	var result httpstat.Result

	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient;
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Server is offline, error:", err)
	}

	if resp.StatusCode != http.StatusOK {
		return false, 0
	}

	dns := int(result.DNSLookup/time.Millisecond)
	tcp := int(result.TCPConnection/time.Millisecond)
	tls := int(result.TLSHandshake/time.Millisecond)
	processing := int(result.ServerProcessing/time.Millisecond)
	totalRequestTime := dns + tcp + tls + processing
	log.Printf("Server response time is: %dms", totalRequestTime)

	return true, totalRequestTime
}

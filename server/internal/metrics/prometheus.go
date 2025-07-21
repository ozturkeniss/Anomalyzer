package metrics

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Prometheus'tan /api/v1/query ile metric çekmek için örnek fonksiyon
func QueryPrometheus(query string) ([]byte, error) {
	url := fmt.Sprintf("http://prometheus:9090/api/v1/query?query=%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

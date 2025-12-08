package ip

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var ip string

func GetIp() string {
	if ip != "" {
		return ip
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	log.Info("Fetching public IP from ip-api.com...")
	resp, err := client.Get("http://ip-api.com/json/?lang=zh-CN")
	if err != nil {
		log.Errorf("Fetch public IP failed: %v", err)
		return "[Your Server IP]"
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			var queryRes map[string]string
			_ = json.Unmarshal(body, &queryRes)
			if val, ok := queryRes["query"]; ok {
				ip = val
				log.Infof("Public IP fetched: %s", ip)
				return ip
			}
		}
	}
	log.Warn("Failed to parse public IP response")
	return "[Your Server IP]"
}

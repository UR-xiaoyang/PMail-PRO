package setup

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jinnrry/pmail/config"

	"github.com/Jinnrry/pmail/i18n"
	"github.com/Jinnrry/pmail/services/auth"
	"github.com/Jinnrry/pmail/utils/context"
	"github.com/Jinnrry/pmail/utils/errors"
	"github.com/Jinnrry/pmail/utils/ip"
	log "github.com/sirupsen/logrus"
)

type DNSItem struct {
	Type  string `json:"type"`
	Host  string `json:"host"`
	Value string `json:"value"`
	TTL   int    `json:"ttl"`
	Tips  string `json:"tips"`
}

func GetDNSSettings(ctx *context.Context) (map[string][]*DNSItem, error) {
	configData, err := config.ReadConfig()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	log.Infof("GetDNSSettings: Read Domains: %v", configData.Domains)

	ipVal := ip.GetIp()
	if ipVal == "[Your Server IP]" {
		if host := ctx.GetValue("Host"); host != nil {
			if hostStr, ok := host.(string); ok && hostStr != "" {
				if h, _, err := net.SplitHostPort(hostStr); err == nil {
					ipVal = h
				} else {
					ipVal = hostStr
				}
			}
		}
	}

	ret := make(map[string][]*DNSItem)

	for _, domain := range configData.Domains {
		ret[domain] = []*DNSItem{
			{Type: "A", Host: strings.ReplaceAll(configData.WebDomain, "."+configData.Domain, ""), Value: ipVal, TTL: 3600, Tips: i18n.GetText(ctx.Lang, "ip_taps")},
			{Type: "A", Host: "smtp", Value: ipVal, TTL: 3600, Tips: i18n.GetText(ctx.Lang, "ip_taps")},
			{Type: "A", Host: "imap", Value: ipVal, TTL: 3600, Tips: i18n.GetText(ctx.Lang, "ip_taps")},
			{Type: "A", Host: "pop", Value: ipVal, TTL: 3600, Tips: i18n.GetText(ctx.Lang, "ip_taps")},
			{Type: "A", Host: "@", Value: ipVal, TTL: 3600, Tips: i18n.GetText(ctx.Lang, "ip_taps")},
			{Type: "MX", Host: "@", Value: fmt.Sprintf("smtp.%s", domain), TTL: 3600},
			{Type: "TXT", Host: "@", Value: "v=spf1 a mx ~all", TTL: 3600},
			{Type: "TXT", Host: "default._domainkey", Value: auth.DkimGen(), TTL: 3600},
		}
	}

	return ret, nil
}

package ddns

import "main/config"

func DoDDNS(ip string, cfconfig config.Cloudflare_Config) {
	if cfconfig.Cloudflare_API_Token != "" && cfconfig.Cloudflare_ZONEID != "" && cfconfig.Cloudflare_Host != "" {

	}
}

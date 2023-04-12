package main

import (
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"log"
)

func main() {
	domain := "status.kncloud.top"
	authEmail := "zw6979014@gmail.com"
	authKey := "76501800a930d81b51ea5e554ebb1ce3bcea0"

	// 创建 Cloudflare DNS Provider 客户端
	config := cloudflare.NewDefaultConfig()
	config.AuthEmail = authEmail
	config.AuthKey = authKey

	client, err := cloudflare.NewDNSProviderConfig(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = client.Present(domain, "_acme-challenge."+domain+".", authKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		// cleanup the challenge after obtaining the certificate
		_ = client.CleanUp(domain, "_acme-challenge."+domain+".", authKey)
	}()
}

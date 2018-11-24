package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var httpsEnvVar = "HEALTHCHECK_HTTPS"

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	httpsEnv, found := os.LookupEnv(httpsEnvVar)
	https := ""
	if !found || len(httpsEnv) == 0 || strings.ToUpper(httpsEnv)[0] != 't' {
		https = "s"
	}
	resp, err := http.Get(fmt.Sprintf("http%s://127.0.0.1/health", https))
	if err != nil {
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		os.Exit(1)
	}

	os.Exit(0)
}

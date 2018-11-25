package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
)

var httpsEnvVar = "HEALTHCHECK_HTTPS"
var portEnvVar = "HEALTHCHECK_PORT"

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	port, _ := os.LookupEnv(portEnvVar)
	if len(port) > 0 {
		port = fmt.Sprintf(":%s", port)
	}
	httpsEnv, found := os.LookupEnv(httpsEnvVar)
	https := ""
	if found && len(httpsEnv) != 0 && strings.ToUpper(httpsEnv)[0] == 'T' {
		https = "s"
	}
	finalUrl := fmt.Sprintf(
		"http%s://127.0.0.1%s/health", https, port,
	)
	resp, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println("Error while connecting:", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		os.Exit(1)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(bytes))

	os.Exit(0)
}

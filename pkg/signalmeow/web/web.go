package web

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const proxyUrlStr = "http://localhost:8080"
const caCertPath = "/Users/sweber/.mitmproxy/mitmproxy-ca-cert.pem"

const urlHost = "chat.signal.org"

// TODO: embed Signal's self-signed cert, and turn off InsecureSkipVerify
func proxiedHTTPClient() *http.Client {
	var proxyURL *url.URL
	if proxyUrlStr != "" {
		var err error
		proxyURL, err = url.Parse(proxyUrlStr)
		if err != nil {
			log.Fatal("Error parsing proxy URL:", err)
		}
	}

	tlsConfig := &tls.Config{}
	if caCertPath != "" {
		var caCert []byte
		var err error
		caCert, err = ioutil.ReadFile(caCertPath)
		if err != nil {
			log.Fatal("Error reading CA certificate:", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig.InsecureSkipVerify = true
		tlsConfig.RootCAs = caCertPool
	}

	transport := &http.Transport{}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	transport.TLSClientConfig = tlsConfig

	client := &http.Client{
		Transport: transport,
	}
	return client
}

func SendHTTPRequest(method string, path string, body []byte, username *string, password *string) (*http.Response, error) {
	urlStr := "https://" + urlHost + path
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("User-Agent", "SignalBridge/0.1")
	//req.Header.Set("X-Signal-Agent", "SignalBridge/0.1")
	if username != nil && password != nil {
		req.SetBasicAuth(*username, *password)
	}

	client := proxiedHTTPClient()
	resp, err := client.Do(req)
	return resp, err
}
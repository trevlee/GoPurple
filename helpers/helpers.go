package helpers

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchUrl(url string) []byte {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // ignore self-signed TLS cert
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)

	_ = res.Body.Close()

	return data
}

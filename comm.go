package espapi

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const endpoint = "https://esp-sandbox.api.gettyimages.com/esp/"

func Response(clientKey string, clientSecret string, espUsername string, espPassword string) (map[string]string, error) {
	v := url.Values{}
	v.Set("client_id", clientKey)
	v.Set("client_secret", clientSecret)
	v.Set("username", espUsername)
	v.Set("password", espPassword)
	v.Set("grant_type", "baz")
	uri := endpoint + "submission/v1/submission_batches"
	log.Infof(uri)
	log.Info(v.Encode())
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.PostForm(uri, v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var record map[string]string
	err = json.Unmarshal(body, &record)
	if err != nil {
		return nil, err
	}
	return record, err
}

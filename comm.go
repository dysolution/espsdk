package espapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const endpoint = "https://esp-sandbox.api.gettyimages.com/esp/"

type ApiClient interface {
	PostBatch(SubmissionBatch) error
	PostRelease(Release) error
}

type Credentials struct {
	ApiKey      string
	ApiSecret   string
	EspUsername string
	EspPassword string
}

type Client struct {
	Credentials
}

func (client Client) PostBatch(b []byte) {
	fmt.Printf("%s\n", b)
	client.Call()
}

func (client Client) PostRelease(r []byte) {
	fmt.Printf("%s\n", r)
	client.Call()
}

func (c Client) Call() {
	v := url.Values{}
	v.Set("client_id", c.ApiKey)
	v.Set("client_secret", c.ApiSecret)
	v.Set("username", c.EspUsername)
	v.Set("password", c.EspPassword)
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

	payload, err := ioutil.ReadAll(resp.Body)
	log.Infof("%s", payload)

	var record map[string]string
	err = json.Unmarshal(payload, &record)
	if errMsg := record["Error"]; errMsg != "" {
		log.Errorf(errMsg)
	}
}

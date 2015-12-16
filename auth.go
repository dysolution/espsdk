package espapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const oauthEndpoint = "https://api.gettyimages.com/oauth2/token"

func credentialsAreIncomplete(c Client) bool {
	return len(c.APIKey) < 1 || len(c.APISecret) < 1 || len(c.ESPUsername) < 1 || len(c.ESPPassword) < 1
}

func (c Client) GetToken() Token {
	if credentialsAreIncomplete(c) {
		log.Fatal("Not all required credentials were supplied.")
	}
	uri := oauthEndpoint
	v := url.Values{}
	v.Set("client_id", c.APIKey)
	v.Set("client_secret", c.APISecret)
	v.Set("username", c.ESPUsername)
	v.Set("password", c.ESPPassword)
	v.Set("grant_type", "password")

	log.Debugf(uri)
	log.Debugf(v.Encode())

	resp, err := http.PostForm(uri, v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	log.Debugf("HTTP %d", resp.StatusCode)
	var token Token
	log.Debugf("%s", payload)

	var response map[string]string
	json.Unmarshal(payload, &response)
	token = Token(response["access_token"])
	log.Debugf("%s", token)
	return token
}

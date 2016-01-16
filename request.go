package espsdk

import (
	"bytes"
	"net/http"
)

// A Request represents the specific API endpoint and action to take. The Object is optional and applies only to endpoints that create or update items (POST and PUT).
type request struct {
	Verb        string `json:"method"`
	Path        string `json:"path"`
	Token       Token  `json:"-"`
	Object      []byte `json:"object"`
	httpRequest *http.Request
}

func newRequest(verb string, path string, token Token, object []byte) *request {
	req, err := http.NewRequest(verb, path, bytes.NewBuffer(object))
	if err != nil {
		log.Fatal(err)
	}
	r := new(request)
	r.Verb = verb
	r.Path = path
	r.Token = token
	r.Object = object
	r.httpRequest = req
	return r
}

func (p *request) requiresAnObject() bool {
	if p.Verb == "POST" || p.Verb == "PUT" || p.Verb == "DELETE" {
		return true
	}
	return false
}

func (p *request) addHeaders(token Token, apiKey string) {
	p.httpRequest.Header.Set("Authorization", "Token token="+string(token))
	p.httpRequest.Header.Set("Content-Type", "application/json")
	p.httpRequest.Header.Set("Api-Key", apiKey)
}

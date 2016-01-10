package espsdk

import "net/url"

// Credentials represent a specific authorized application performing
// operations on objects belonging to a specific ESP user.
type credentials struct {
	APIKey      string
	APISecret   string
	ESPUsername string
	ESPPassword string
}

func (c *credentials) areInvalid() bool {
	if len(c.APIKey) < 1 || len(c.APISecret) < 1 || len(c.ESPUsername) < 1 || len(c.ESPPassword) < 1 {
		return true
	}
	return false
}

func formValues(c *credentials) url.Values {
	v := url.Values{}
	v.Set("client_id", c.APIKey)
	v.Set("client_secret", c.APISecret)
	v.Set("username", c.ESPUsername)
	v.Set("password", c.ESPPassword)
	v.Set("grant_type", "password")
	return v
}

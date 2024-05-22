package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) login(payload *KLoginPayload) (*KLoginRes, error) {

	formData := url.Values{
		"client_id":     {payload.clientId},
		"client_secret": {payload.clientSecret},
		"grant_type":    {payload.grantType},
		"username":      {payload.username},
		"password":      {payload.password},
	}

	encodedFormData := formData.Encode()

	req, err := http.NewRequest("POST", "http://localhost:8181/realms/PSKE/protocol/openid-connect/token", strings.NewReader(encodedFormData))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Login was unsuccessful")
	}

	kLoginRes := &KLoginRes{}

	json.NewDecoder(resp.Body).Decode(kLoginRes)

	return kLoginRes, nil
}

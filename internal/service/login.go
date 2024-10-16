package service

import (
	"costa-wifi/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	BasicAuth = "Basic Z2lneTo="
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(costaNumber, birthDate string) (*LoginResponse, error) {
	apiURL := "https://mobileapp.api.costa.it/oauth/token"

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("user", "gigy")
	data.Set("birthDate", birthDate)
	data.Set("costaNumber", costaNumber)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", BasicAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("DNT", "1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	err = config.WriteConfigValue(config.KeyJWTToken, loginResp.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("error saving access token to config: %w", err)
	}

	return &loginResp, nil
}

func GetToken() (string, error) {
    token, err := config.ReadConfigValue(config.KeyJWTToken)
    if err != nil {
        return "", fmt.Errorf("error reading token from config: %w", err)
    }

    expired, err := IsTokenExpired(token)
    if err != nil {
        return "", fmt.Errorf("error checking token expiration: %w", err)
    }

    if expired {
        cardID, err := config.ReadConfigValue(config.KeyCard)
        if err != nil {
            return "", fmt.Errorf("error reading card ID from config: %w", err)
        }

        birthday, err := config.ReadConfigValue(config.KeyBithday)
        if err != nil {
            return "", fmt.Errorf("error reading birthday from config: %w", err)
        }

        loginResp, err := Login(cardID, birthday)
        if err != nil {
            return "", fmt.Errorf("error refreshing token: %w", err)
        }

        err = config.WriteConfigValue(config.KeyJWTToken, loginResp.AccessToken)
        if err != nil {
            return "", fmt.Errorf("error saving new token to config: %w", err)
        }

        return loginResp.AccessToken, nil
    }

    return token, nil
}

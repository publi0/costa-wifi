package service

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type ConnectRequest struct {
    BookingID string `json:"bookingId"`
    IPAddress string `json:"ipAddress"`
}

type SessionData struct {
    SessionID  string `json:"sessionId"`
    StartTime  string `json:"startTime"`
    IPAddress  string `json:"ipAddress"`
    MacAddress string `json:"macAddress"`
    Status     string `json:"status"`
    UserAgent  string `json:"userAgent"`
    Account    string `json:"account"`
}

type ConnectResponse struct {
    Data                    SessionData         `json:"data"`
    Success                 bool                `json:"success"`
    Error                   string              `json:"error"`
    Code                    string              `json:"code"`
    AdditionalInformations  map[string]interface{} `json:"additionalInformations"`
}

func ConnectSession(bookingID, ipAddress string) (*ConnectResponse, error) {
    token, err := GetToken()
    if err != nil {
        return nil, fmt.Errorf("error getting token: %w", err)
    }

    reqBody := ConnectRequest{
        BookingID: bookingID,
        IPAddress: ipAddress,
    }

    jsonBody, err := json.Marshal(reqBody)
    if err != nil {
        return nil, fmt.Errorf("error marshaling request body: %w", err)
    }

    req, err := http.NewRequest("POST", "https://mobileapp.api.costa.it/api/ipackages/v2/sessions", bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }

    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Accept-Language", "en-US,pt-BR;q=0.7,en;q=0.3")
    req.Header.Set("Cache-Control", "no-cache")
    req.Header.Set("Pragma", "no-cache")
    req.Header.Set("Origin", "https://mobileapp.aem.costa.it")
    req.Header.Set("Referer", "https://mobileapp.aem.costa.it/")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:130.0) Gecko/20100101 Firefox/130.0")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var connectResp ConnectResponse
    err = json.NewDecoder(resp.Body).Decode(&connectResp)
    if err != nil {
        return nil, fmt.Errorf("error decoding response: %w", err)
    }

    return &connectResp, nil
}

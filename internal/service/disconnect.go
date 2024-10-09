package service

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type DisconnectRequest struct {
    SessionID string `json:"sessionId"`
    Status    string `json:"status"`
}

func DisconnectSession(sessionID string) error {
    token, err := GetToken()
    if err != nil {
        return fmt.Errorf("error getting token: %w", err)
    }

    reqBody := DisconnectRequest{
        SessionID: sessionID,
        Status:    "INACTIVE",
    }

    jsonBody, err := json.Marshal(reqBody)
    if err != nil {
        return fmt.Errorf("error marshaling request body: %w", err)
    }

    req, err := http.NewRequest("POST", "https://mobileapp.api.costa.it/api/ipackages/v2/sessions", bytes.NewBuffer(jsonBody))
    if err != nil {
        return fmt.Errorf("error creating request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}

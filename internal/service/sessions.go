package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type SessionsRequest struct {
	Account         string   `json:"account"`
	IncludeSessions bool     `json:"includeSessions"`
	IncludeUpgrades bool     `json:"includeUpgrades"`
	Lang            string   `json:"lang"`
	ProductTypes    []string `json:"productTypes"`
	Status          []string `json:"status"`
}

type SessionUser struct {
	SessionID  string `json:"sessionId"`
	StartTime  string `json:"startTime"`
	IPAddress  string `json:"ipAddress"`
	MacAddress string `json:"macAddress"`
	Status     string `json:"status"`
	UserAgent  string `json:"userAgent"`
	Account    string `json:"account"`
}

type SessionsResponse struct {
	Data []struct {
		FreeSites []struct {
			Link  string `json:"link"`
			Title string `json:"title"`
		} `json:"freeSites"`
		InternetPackagesCategories []struct {
			LongDescription  string `json:"longDescription"`
			ImageDetail      string `json:"imageDetail"`
			Code             string `json:"code"`
			Name             string `json:"name"`
			ID               string `json:"id"`
			ShortDescription string `json:"shortDescription"`
			InternetPackages []struct {
				Code           string `json:"code"`
				ID             string `json:"id"`
				Category       string `json:"category"`
				Title          string `json:"title"`
				Version        int64  `json:"version"`
				PackageDetails struct {
					BookingID string `json:"bookingId"`
					Sessions []SessionUser `json:"sessions"`
				} `json:"packageDetails"`
			} `json:"internetPackages"`
		} `json:"internetPackagesCategories"`
	} `json:"data"`
}

func GetPlanSessions() (*SessionsResponse, error) {
	token, err := GetToken()
	if err != nil {
		return nil, fmt.Errorf("error getting token: %w", err)
	}
	account, err := ExtractGuestID(token)
	if err != nil {
		return nil, fmt.Errorf("error getting account id: %w", err)
	}

	account = extractNumbers(account)

	reqBody := SessionsRequest{
		Account:         account,
		IncludeSessions: true,
		IncludeUpgrades: false,
		Lang:            "en",
		ProductTypes:    []string{"INTERNET"},
		Status:          []string{"ACTIVE", "INACTIVE"},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", "https://mobileapp.api.costa.it/api/ipackages/v2/bookings-list", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var sessionsResp SessionsResponse
	err = json.NewDecoder(resp.Body).Decode(&sessionsResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &sessionsResp, nil
}

func GetUserSession(sessionsResp *SessionsResponse) ([]SessionUser, error) {
    if sessionsResp == nil {
        return nil, fmt.Errorf("sessions response is nil")
    }

    if len(sessionsResp.Data) == 0 {
        return nil, fmt.Errorf("no data in sessions response")
    }

    if len(sessionsResp.Data[0].InternetPackagesCategories) == 0 {
        return nil, fmt.Errorf("no internet packages categories found")
    }

    if len(sessionsResp.Data[0].InternetPackagesCategories[0].InternetPackages) == 0 {
        return nil, fmt.Errorf("no internet packages found")
    }

    sessions := sessionsResp.Data[0].InternetPackagesCategories[0].InternetPackages[0].PackageDetails.Sessions

    if len(sessions) == 0 {
        return nil, fmt.Errorf("no sessions found")
    }

    return sessions, nil
}

func extractNumbers(s string) string {
	re := regexp.MustCompile(`\d+`)
	return re.FindString(s)
}

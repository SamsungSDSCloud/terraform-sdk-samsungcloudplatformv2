package scpsdk

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type Endpoint struct {
	Region      string `json:"region"`
	ServiceType string `json:"service_type"`
	ServiceName string `json:"service_name"`
	URL         string `json:"url"`
}

type Catalog struct {
	AuthURL       string
	AccessKey     string
	SecretKey     string
	DefaultRegion string
}

var catalog []Endpoint

func NewCatalog(authURL, accessKey, secretKey, defaultRegion string) *Catalog {
	return &Catalog{
		AuthURL:       authURL,
		AccessKey:     accessKey,
		SecretKey:     secretKey,
		DefaultRegion: defaultRegion,
	}
}

func (c *Catalog) GetEndpoint(serviceType, region, accountID string) (string, error) {
	if len(catalog) == 0 {
		req, err := http.NewRequest("GET", c.AuthURL+"/endpoints", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return "", err
		}

		timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		signature := MakeHmacSignature("GET", c.AuthURL+"/endpoints", timestamp, c.AccessKey, c.SecretKey, accountID, ClientType)

		req.Header.Set(HeaderClientType, ClientType)
		req.Header.Set(HeaderLanguage, "en-US")
		req.Header.Set(HeaderAccountId, accountID)
		req.Header.Set(HeaderAccessKey, c.AccessKey)
		req.Header.Set(HeaderTimestamp, timestamp)
		req.Header.Set(HeaderSignature, signature)

		httpClient := &http.Client{
			Transport: &http.Transport{
				// Disable SSL verification (insecure)
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return "", err
		}

		body, err := io.ReadAll(resp.Body)

		var response struct {
			Endpoints []Endpoint `json:"endpoints"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return "", err
		}

		for _, endpoint := range response.Endpoints {
			catalog = append(catalog, endpoint)
		}
	}

	var candidates []Endpoint

	for _, endpoint := range catalog {
		if endpoint.ServiceType != serviceType {
			continue
		}
		candidates = append(candidates, endpoint)
	}

	if len(candidates) == 0 {
		return "", errors.New("no matching endpoint found")
	} else if len(candidates) == 1 {
		return candidates[0].URL, nil
	}

	originalCandidates := make([]Endpoint, len(candidates))
	copy(originalCandidates, candidates)
	candidates = []Endpoint{}

	if region == "" {
		region = c.DefaultRegion
	}

	if region != "" {
		for _, endpoint := range originalCandidates {
			if endpoint.Region != region {
				continue
			}
			candidates = append(candidates, endpoint)
		}
	}

	if len(candidates) == 0 {
		candidates = originalCandidates
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].URL < originalCandidates[j].URL
	})

	return candidates[0].URL, nil
}

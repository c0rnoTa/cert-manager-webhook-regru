package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.reg.ru/api/regru2"
)

type RegruClient struct {
	username string
	password string
	zone     string
}

func NewRegruClient(username string, password string, zone string) *RegruClient {
	return &RegruClient{
		username: username,
		password: password,
		zone:     zone,
	}
}

func (c *RegruClient) getRecords() error {
	inputData := fmt.Sprintf(`{"username":"%s","password":"%s","domains":[{"dname":"%s"}],"output_content_type":"plain"}`, c.username, c.password, c.zone)

	requestData := url.Values{}
	requestData.Set("input_data", inputData)
	requestData.Set("input_format", "json")

	requestURL := fmt.Sprintf("%s/zone/get_resource_records", defaultBaseURL)

	fmt.Printf("[%s] Get resource records\n", c.zone)
	res, err := http.PostForm(requestURL, requestData)
	if err != nil {
		return fmt.Errorf("[%s] failed to make get resource records request POST %s: %v", c.zone, requestURL, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("[%s] failed to read get resource records response %d body: %v", c.zone, res.StatusCode, err)
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("[%s] get resource records response failed with status code %d and body: %s", c.zone, res.StatusCode, body)
	}

	fmt.Printf("[%s] Get TXT resource record success. Response body: %s\n", c.zone, body)

	return nil
}

func (c *RegruClient) createTXT(domain string, value string) error {
	inputData := fmt.Sprintf(`{"username":"%s","password":"%s","domains":[{"dname":"%s"}],"subdomain":"%s","text":"%s","output_content_type":"plain"}`, c.username, c.password, c.zone, domain, value)

	requestData := url.Values{}
	requestData.Set("input_data", inputData)
	requestData.Set("input_format", "json")

	requestURL := fmt.Sprintf("%s/zone/add_txt", defaultBaseURL)

	fmt.Printf("[%s] Add TXT record\n", c.zone)
	res, err := http.PostForm(requestURL, requestData)
	if err != nil {
		return fmt.Errorf("[%s] failed to make add TXT record request POST %s: %v", c.zone, requestURL, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("[%s] failed to read add TXT record response %d body: %v", c.zone, res.StatusCode, err)
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("[%s] add TXT record response failed with status code %d and body: %s", c.zone, res.StatusCode, body)
	}

	fmt.Printf("[%s] Add TXT record success. Response body: %s\n", c.zone, body)

	return nil
}

func (c *RegruClient) deleteTXT(domain string, value string) error {
	inputData := fmt.Sprintf(`{"username":"%s","password":"%s","domains":[{"dname":"%s"}],"subdomain":"%s","content":"%s","record_type":"TXT","output_content_type":"plain"}`, c.username, c.password, c.zone, domain, value)

	requestData := url.Values{}
	requestData.Set("input_data", inputData)
	requestData.Set("input_format", "json")

	requestURL := fmt.Sprintf("%s/zone/remove_record", defaultBaseURL)

	fmt.Printf("[%s] Remove TXT record\n", c.zone)
	res, err := http.PostForm(requestURL, requestData)
	if err != nil {
		return fmt.Errorf("[%s] failed to make remove TXT record request POST %s: %v", c.zone, requestURL, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("[%s] failed to read remove TXT record response %d body: %v", c.zone, res.StatusCode, err)
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("[%s] remove TXT record response failed with status code %d and body: %s", c.zone, res.StatusCode, body)
	}

	fmt.Printf("[%s] Remove TXT record success. Response body: %s\n", c.zone, body)

	return nil
}

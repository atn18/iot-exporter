package iot

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	config *Config
}

func NewClient(config *Config) *Client {
	return &Client{config: config}
}

func (c *Client) GetIOTInfo() (*IOTInfo, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.Host+"/v1.0/user/info", nil)
	if err != nil {
		log.Print("Could not create new request")
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", c.config.Token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error occurred while requesting devices info")
		return nil, err
	}

	var dataResp = &IOTInfo{}
	err = json.NewDecoder(resp.Body).Decode(dataResp)
	if err != nil {
		log.Print("Error occurred while decoding response body")
		return nil, err
	}

	if dataResp.Status != "ok" {
		log.Printf("Request has completed with error status. Message: %s", dataResp.Message)
		return nil, errors.New("request has completed with error status")
	}

	return dataResp, nil
}

func (c *Client) GetDeviceInfo(deviceId string) (*Device, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.Host+"/v1.0/devices/"+deviceId, nil)
	if err != nil {
		log.Print("Could not create new request")
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", c.config.Token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error occurred while requesting device info")
		return nil, err
	}

	var dataResp = &Device{}
	err = json.NewDecoder(resp.Body).Decode(dataResp)
	if err != nil {
		log.Print("Error occurred while decoding response body")
		return nil, err
	}

	if dataResp.Status != "ok" {
		log.Printf("Request has completed with error status. Message: %s", dataResp.Message)
		return nil, errors.New("request has completed with error status")
	}

	return dataResp, nil
}

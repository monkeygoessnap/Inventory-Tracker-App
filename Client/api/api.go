/*
Package API provides the functions required to send a request to
the API server, and parse the JSON response into local variables
to work with.
*/
package api

import (
	"PGL/Client/auth"
	"PGL/Client/log"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Base url of the APIServer
const baseURL = "https://localhost:8080/api/v1"

//Get all
func GetAll(model, idparam string) []byte {
	//generate the JWT token based on the secret key to authenticate the client
	token, _ := auth.GenerateJWT()
	url := fmt.Sprintf("%s/%s?userid=%s", baseURL, model, idparam)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Token", token)
	//skip secure auth for development env
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(req)
	if err != nil {
		log.Warning.Println(err)
		return nil
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return data
	}
}

//Get function
func Get(model, params string) []byte {

	token, _ := auth.GenerateJWT()
	url := fmt.Sprintf("%s/%s/%s", baseURL, model, params)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Token", token)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(req)
	if err != nil {
		log.Warning.Println(err)
		return nil
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return data
	}
}

//Add function
func Add(model string, jsonData interface{}) []byte {

	token, _ := auth.GenerateJWT()
	jsonValue, _ := json.Marshal(jsonData)
	url := fmt.Sprintf("%s/%s", baseURL, model)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Token", token)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(req)
	if err != nil {
		log.Warning.Println(err)
		return nil
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return data
	}
}

//Edit function
func Edit(model, params string, jsonData interface{}) []byte {

	token, _ := auth.GenerateJWT()
	jsonValue, _ := json.Marshal(jsonData)
	url := fmt.Sprintf("%s/%s/%s", baseURL, model, params)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Token", token)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(req)
	if err != nil {
		log.Warning.Println(err)
		return nil
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return data
	}
}

//Delete function
func Del(model, params string) []byte {

	token, _ := auth.GenerateJWT()
	url := fmt.Sprintf("%s/%s/%s", baseURL, model, params)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Token", token)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(req)
	if err != nil {
		log.Warning.Println(err)
		return nil
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return data
	}
}

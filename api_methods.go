package op5

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (o Payload) Get(server Server) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	body := bytes.NewReader(o.Body)
	req, err := http.NewRequest("GET", "https://"+server.Host+"/api/config/"+o.Endpoint+"/"+o.Target, body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if server.AllowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	e, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	status := response.Status
	code := response.StatusCode
	err = json.Unmarshal(e, &returnedError)
	return e, status, code, err
}

func (o Payload) Add(server Server) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	body := bytes.NewReader(o.Body)
	fmt.Println(server.Host)
	req, err := http.NewRequest("POST", "https://"+ server.Host +"/api/config/"+o.Endpoint, body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if server.AllowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	e, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	status := response.Status
	code := response.StatusCode
	err = json.Unmarshal(e, &returnedError)
	return e, status, code, err
}

func (o Payload) Update(server Server) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	body := bytes.NewReader(o.Body)
	ep := fmt.Sprintf("https://%s/api/config/%s/%s", server.Host, o.Endpoint, o.Target)
	req, err := http.NewRequest("PATCH", ep, body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if server.AllowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	e, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	status := response.Status
	code := response.StatusCode
	err = json.Unmarshal(e, &returnedError)
	return e, status, code, err
}

func (o Payload) Replace(server Server) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	body := bytes.NewReader(o.Body)
	ep := fmt.Sprintf("https://%s/api/config/%s/%s", server.Host, o.Endpoint, o.Target)
	req, err := http.NewRequest("PUT", ep, body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if server.AllowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	e, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	status := response.Status
	code := response.StatusCode
	err = json.Unmarshal(e, &returnedError)

	return e, status, code, err
}

func (o Payload) Delete(server Server) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	body := bytes.NewReader(o.Body)
	ep := fmt.Sprintf("https://%s/api/config/%s/%s", server.Host, o.Endpoint, o.Target)
	req, err := http.NewRequest("PUT", ep, body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if server.AllowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	e, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	status := response.Status
	code := response.StatusCode
	err = json.Unmarshal(e, &returnedError)
	return e, status, code, err
}

func (o Payload) Command(server Server) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	body := bytes.NewReader(o.Body)
	req, err := http.NewRequest("GET", "https://"+server.Host+"/api/command/"+o.Target, body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(server.Username, server.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if server.AllowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	e, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	status := response.Status
	code := response.StatusCode
	err = json.Unmarshal(e, &returnedError)
	return e, status, code, err
}

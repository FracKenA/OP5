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

type UserGroup struct {
	AccessRights []string `json:"access_rights,omitempty"`
	Name         string   `json:"name,omitempty"`
}

func GetUserGroup(userGroup string, server string, userGroupname string, password string, allowInsecure bool) (UserGroup, error) {

	var returnedBody UserGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/usergroup/"+userGroup, nil)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(userGroupname, password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if allowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		// You may read / inspect response body
		e, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		err = json.Unmarshal(e, &returnedError)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	err = json.Unmarshal(b, &returnedBody)

	return returnedBody, err
}

func CreateUserGroup(data UserGroup, server string, userGroupname string, password string, allowInsecure bool) (UserGroup, error) {

	var returnedBody UserGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/config/usergroup", body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(userGroupname, password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if allowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		// You may read / inspect response body
		e, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		err = json.Unmarshal(e, &returnedError)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	err = json.Unmarshal(b, &returnedBody)

	return returnedBody, err
}

func UpdateUserGroup(data UserGroup, server string, userGroupname string, password string, allowInsecure bool) (UserGroup, error) {

	var returnedBody UserGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PATCH", "https://"+server+"/api/config/usergroup", body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(userGroupname, password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if allowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		// You may read / inspect response body
		e, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		err = json.Unmarshal(e, &returnedError)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	err = json.Unmarshal(b, &returnedBody)

	return returnedBody, err
}


func ReplaceUserGroup(data UserGroup, server string, userGroupname string, password string, allowInsecure bool) (UserGroup, error) {

	var returnedBody UserGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PUT", "https://"+server+"/api/config/usergroup", body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(userGroupname, password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if allowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		// You may read / inspect response body
		e, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		err = json.Unmarshal(e, &returnedError)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()

	err = json.Unmarshal(b, &returnedBody)

	return returnedBody, err
}

func DeleteUserGroup(userGroup string, server string, userGroupname string, password string, allowInsecure bool) (string, error) {

	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/usergroup/"+userGroup, nil)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(userGroupname, password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if allowInsecure == true {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	} else {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	op5client := http.Client{Transport: customTransport, Timeout: 15 * time.Second}

	response, err := op5client.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		// You may read / inspect response body
		e, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		err = json.Unmarshal(e, &returnedError)
	}


	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()
	body, _ := ioutil.ReadAll(response.Body)

	return string(body), err
}

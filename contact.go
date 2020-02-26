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

type Contact struct {
	ContactName                 string        `json:"contact_name,omitempty"`
	Alias                       string        `json:"alias,omitempty"`
	HostNotificationOptions     []string      `json:"host_notification_options,omitempty"`
	ServiceNotificationOptions  []string      `json:"service_notification_options,omitempty"`
	HostNotificationCmds        string        `json:"host_notification_cmds,omitempty"`
	ServiceNotificationCmds     string        `json:"service_notification_cmds,omitempty"`
	Register                    bool          `json:"register,omitempty"`
	Template                    string        `json:"template,omitempty"`
	FileID                      string        `json:"file_id,omitempty"`
	HostNotificationsEnabled    bool          `json:"host_notifications_enabled,omitempty"`
	ServiceNotificationsEnabled bool          `json:"service_notifications_enabled,omitempty"`
	CanSubmitCommands           bool          `json:"can_submit_commands,omitempty"`
	RetainStatusInformation     bool          `json:"retain_status_information,omitempty"`
	RetainNonstatusInformation  bool          `json:"retain_nonstatus_information,omitempty"`
	HostNotificationPeriod      string        `json:"host_notification_period,omitempty"`
	ServiceNotificationPeriod   string        `json:"service_notification_period,omitempty"`
	HostNotificationCmdsArgs    string        `json:"host_notification_cmds_args,omitempty"`
	ServiceNotificationCmdsArgs string        `json:"service_notification_cmds_args,omitempty"`
	Contactgroups               []interface{} `json:"contactgroups,omitempty"`
	Email                       string        `json:"email,omitempty"`
	Pager                       string        `json:"pager,omitempty"`
	Address1                    string        `json:"address1,omitempty"`
	Address2                    string        `json:"address2,omitempty"`
	Address3                    string        `json:"address3,omitempty"`
	Address4                    string        `json:"address4,omitempty"`
	Address5                    string        `json:"address5,omitempty"`
	Address6                    string        `json:"address6,omitempty"`
	EnableAccess                string        `json:"enable_access,omitempty"`
	Realname                    string        `json:"realname,omitempty"`
	Password                    string        `json:"password,omitempty"`
	PasswordAlgo                string        `json:"password_algo,omitempty"`
	Modules                     []interface{} `json:"modules,omitempty"`
	Groups                      []interface{} `json:"groups,omitempty"`
	AccessLevels                string        `json:"Access Levels,omitempty"`
	RealnameUser                string        `json:"Realname,omitempty"`
	PasswordUser                string        `json:"Password,omitempty"`
	Role                        []interface{} `json:"Role,omitempty"`
}

func GetContact(contact string, server string, username string, password string, allowInsecure bool) (Contact, error) {

	var returnedBody Contact
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/contact/"+contact, nil)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(username, password)
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

func CreateContact(data Contact, server string, username string, password string, allowInsecure bool) (Contact, error) {

	var returnedBody Contact
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/config/contact", body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(username, password)
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

func UpdateContact(data Contact, server string, username string, password string, allowInsecure bool) (Contact, error) {

	var returnedBody Contact
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PATCH", "https://"+server+"/api/config/contact", body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(username, password)
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


func ReplaceContact(data Contact, server string, username string, password string, allowInsecure bool) (Contact, error) {

	var returnedBody Contact
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PUT", "https://"+server+"/api/config/contact", body)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(username, password)
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

func DeleteContact(contact string, server string, username string, password string, allowInsecure bool) (string, error) {

	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/contact/"+contact, nil)
	if err != nil {
		fmt.Printf("Error%s\n", err)

	}
	req.SetBasicAuth(username, password)
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

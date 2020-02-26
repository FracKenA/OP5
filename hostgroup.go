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

type HostGroup struct {
	HostgroupName    string        `json:"hostgroup_name,omitempty"`
	Alias            string        `json:"alias,omitempty"`
	Register         bool          `json:"register,omitempty"`
	FileID           string        `json:"file_id,omitempty"`
	Members          []interface{} `json:"members,omitempty"`
	HostgroupMembers []interface{} `json:"hostgroup_members,omitempty"`
	Notes            string        `json:"notes,omitempty"`
	NotesURL         string        `json:"notes_url,omitempty"`
	ActionURL        string        `json:"action_url,omitempty"`
	Services         []struct {
		ServiceDescription         string        `json:"service_description,omitempty"`
		CheckCommand               string        `json:"check_command,omitempty"`
		CheckCommandArgs           string        `json:"check_command_args,omitempty"`
		ParallelizeCheck           bool          `json:"parallelize_check,omitempty"`
		Obsess                     bool          `json:"obsess,omitempty"`
		StalkingOptions            []string      `json:"stalking_options,omitempty"`
		Template                   string        `json:"template,omitempty"`
		Register                   bool          `json:"register,omitempty"`
		FileID                     string        `json:"file_id,omitempty"`
		IsVolatile                 bool          `json:"is_volatile,omitempty"`
		MaxCheckAttempts           int           `json:"max_check_attempts,omitempty"`
		CheckInterval              int           `json:"check_interval,omitempty"`
		RetryInterval              int           `json:"retry_interval,omitempty"`
		ActiveChecksEnabled        bool          `json:"active_checks_enabled,omitempty"`
		PassiveChecksEnabled       bool          `json:"passive_checks_enabled,omitempty"`
		CheckPeriod                string        `json:"check_period,omitempty"`
		EventHandlerEnabled        bool          `json:"event_handler_enabled,omitempty"`
		FlapDetectionEnabled       bool          `json:"flap_detection_enabled,omitempty"`
		ProcessPerfData            bool          `json:"process_perf_data,omitempty"`
		RetainStatusInformation    bool          `json:"retain_status_information,omitempty"`
		RetainNonstatusInformation bool          `json:"retain_nonstatus_information,omitempty"`
		NotificationInterval       int           `json:"notification_interval,omitempty"`
		NotificationPeriod         string        `json:"notification_period,omitempty"`
		NotificationOptions        []string      `json:"notification_options,omitempty"`
		NotificationsEnabled       bool          `json:"notifications_enabled,omitempty"`
		HostName                   string        `json:"host_name,omitempty"`
		DisplayName                string        `json:"display_name,omitempty"`
		Servicegroups              []interface{} `json:"servicegroups,omitempty"`
		CheckFreshness             string        `json:"check_freshness,omitempty"`
		FreshnessThreshold         string        `json:"freshness_threshold,omitempty"`
		EventHandler               string        `json:"event_handler,omitempty"`
		EventHandlerArgs           string        `json:"event_handler_args,omitempty"`
		LowFlapThreshold           string        `json:"low_flap_threshold,omitempty"`
		HighFlapThreshold          string        `json:"high_flap_threshold,omitempty"`
		FlapDetectionOptions       []interface{} `json:"flap_detection_options,omitempty"`
		FirstNotificationDelay     string        `json:"first_notification_delay,omitempty"`
		Contacts                   []interface{} `json:"contacts,omitempty"`
		ContactGroups              []interface{} `json:"contact_groups,omitempty"`
		Notes                      string        `json:"notes,omitempty"`
		NotesURL                   string        `json:"notes_url,omitempty"`
		ActionURL                  string        `json:"action_url,omitempty"`
		IconImage                  string        `json:"icon_image,omitempty"`
		IconImageAlt               string        `json:"icon_image_alt,omitempty"`
		ObsessOverService          bool          `json:"obsess_over_service,omitempty"`
	} `json:"services,omitempty"`
}

func GetHostGroup(hostGroup string, server string, username string, password string, allowInsecure bool) (HostGroup, error) {

	var returnedBody HostGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/hostGroup/"+hostGroup, nil)
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

func CreateHostGroup(data HostGroup, server string, username string, password string, allowInsecure bool) (HostGroup, error) {

	var returnedBody HostGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/config/hostGroup", body)
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

func UpdateHostGroup(data HostGroup, server string, username string, password string, allowInsecure bool) (HostGroup, error) {

	var returnedBody HostGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PATCH", "https://"+server+"/api/config/hostGroup", body)
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


func ReplaceHostGroup(data HostGroup, server string, username string, password string, allowInsecure bool) (HostGroup, error) {

	var returnedBody HostGroup
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PUT", "https://"+server+"/api/config/hostGroup", body)
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

func DeleteHostGroup(hostGroup string, server string, username string, password string, allowInsecure bool) (string, error) {

	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/hostGroup/"+hostGroup, nil)
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

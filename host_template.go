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

type HostTemplate struct {
	CheckCommand               string        `json:"check_command,omitempty"`
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
	Name                       string        `json:"name,omitempty"`
	Register                   bool          `json:"register,omitempty"`
	FileID                     string        `json:"file_id,omitempty"`
	Template                   string        `json:"template,omitempty"`
	Alias                      string        `json:"alias,omitempty"`
	DisplayName                string        `json:"display_name,omitempty"`
	Address                    string        `json:"address,omitempty"`
	Hostgroups                 []interface{} `json:"hostgroups,omitempty"`
	Parents                    []interface{} `json:"parents,omitempty"`
	Children                   []interface{} `json:"children,omitempty"`
	CheckCommandArgs           string        `json:"check_command_args,omitempty"`
	Contacts                   []interface{} `json:"contacts,omitempty"`
	ContactGroups              []interface{} `json:"contact_groups,omitempty"`
	Obsess                     string        `json:"obsess,omitempty"`
	CheckFreshness             string        `json:"check_freshness,omitempty"`
	FreshnessThreshold         string        `json:"freshness_threshold,omitempty"`
	EventHandler               string        `json:"event_handler,omitempty"`
	EventHandlerArgs           string        `json:"event_handler_args,omitempty"`
	LowFlapThreshold           string        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold          string        `json:"high_flap_threshold,omitempty"`
	FlapDetectionOptions       []interface{} `json:"flap_detection_options,omitempty"`
	FirstNotificationDelay     string        `json:"first_notification_delay,omitempty"`
	StalkingOptions            []interface{} `json:"stalking_options,omitempty"`
	IconImage                  string        `json:"icon_image,omitempty"`
	IconImageAlt               string        `json:"icon_image_alt,omitempty"`
	StatusmapImage             string        `json:"statusmap_image,omitempty"`
	Notes                      string        `json:"notes,omitempty"`
	ActionURL                  string        `json:"action_url,omitempty"`
	NotesURL                   string        `json:"notes_url,omitempty"`
	TwoDCoords                 string        `json:"2d_coords,omitempty"`
	ObsessOverHost             string        `json:"obsess_over_host,omitempty"`
}

func GetHostTemplate(hostTemplate string, server string, username string, password string, allowInsecure bool) (HostTemplate, error) {

	var returnedBody HostTemplate
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/host_template/"+hostTemplate, nil)
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

func CreateHostTemplate(data HostTemplate, server string, username string, password string, allowInsecure bool) (HostTemplate, error) {

	var returnedBody HostTemplate
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/config/host_template", body)
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

func UpdateHostTemplate(data HostTemplate, server string, username string, password string, allowInsecure bool) (HostTemplate, error) {

	var returnedBody HostTemplate
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PATCH", "https://"+server+"/api/config/host_template", body)
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


func ReplaceHostTemplate(data HostTemplate, server string, username string, password string, allowInsecure bool) (HostTemplate, error) {

	var returnedBody HostTemplate
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PUT", "https://"+server+"/api/config/host_template", body)
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

func DeleteHostTemplate(hostTemplate string, server string, username string, password string, allowInsecure bool) (string, error) {

	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	req, err := http.NewRequest("DELETE", "https://"+server+"/api/config/host_template/"+hostTemplate, nil)
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

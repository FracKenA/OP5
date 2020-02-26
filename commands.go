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

type AckHostProblem struct {
	HostName   string `json:"host_name"`
	Sticky     int    `json:"sticky"`
	Notify     bool   `json:"notify"`
	Persistent bool   `json:"persistent"`
	Comment    string `json:"comment"`
}

type AckServiceProblem struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	Sticky             int    `json:"sticky"`
	Notify             bool   `json:"notify"`
	Persistent         bool   `json:"persistent"`
	Comment            string `json:"comment"`
}

type DelDowntimeByHostName struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	DowntimeStartTime  string `json:"downtime_start_time"`
	Comment            string `json:"comment"`
}

type ProcHostCheckResult struct {
	HostName     string `json:"host_name"`
	StatusCode   int    `json:"status_code"`
	PluginOutput string `json:"plugin_output"`
}

type ProcServiceCheckResult struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	StatusCode         int    `json:"status_code"`
	PluginOutput       string `json:"plugin_output"`
}

type SchPropagateHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     bool   `json:"fixed"`
	TriggerID int    `json:"trigger_id"`
	Duration  int    `json:"duration"`
	Comment   string `json:"comment"`
}

type SchPropagateTriggeredHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     bool   `json:"fixed"`
	TriggerID int    `json:"trigger_id"`
	Duration  int    `json:"duration"`
	Comment   string `json:"comment"`
}

type SchHostCheck struct {
	HostName  string `json:"host_name"`
	CheckTime string `json:"check_time"`
}

type SchHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     bool   `json:"fixed"`
	TriggerID int    `json:"trigger_id"`
	Duration  int    `json:"duration"`
	Comment   string `json:"comment"`
}

type SchServiceCheck struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	CheckTime          string `json:"check_time"`
}

type SchServiceDowntime struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	StartTime          string `json:"start_time"`
	EndTime            string `json:"end_time"`
	Fixed              bool   `json:"fixed"`
	TriggerID          int    `json:"trigger_id"`
	Duration           int    `json:"duration"`
	Comment            string `json:"comment"`
}

func AcknowledgeHostProblem(data AckHostProblem, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/ACKNOWLEDGE_HOST_PROBLEM", body)
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
		fmt.Println(string(e))
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

func AcknowledgeServiceProblem(data AckServiceProblem, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/ACKNOWLEDGE_SVC_PROBLEM", body)
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
		fmt.Println(string(e))
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

func DeleteDowntimeByHostName(data DelDowntimeByHostName, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/DEL_DOWNTIME_BY_HOST_NAME", body)
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
		fmt.Println(string(e))
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

func ProcessHostCheckResult(data ProcHostCheckResult, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/PROCESS_HOST_CHECK_RESULT", body)
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
		fmt.Println(string(e))
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

func ProcessServiceCheckResult(data ProcServiceCheckResult, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/PROCESS_SERVICE_CHECK_RESULT", body)
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
		fmt.Println(string(e))
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

func ScheduleAndPropagateHostDowntime(data SchPropagateHostDowntime, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME", body)
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
		fmt.Println(string(e))
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

func ScheduleAndPropagateTriggeredHostDowntime(data SchPropagateTriggeredHostDowntime, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME", body)
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
		fmt.Println(string(e))
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

func ScheduleHostCheck(data SchHostCheck, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/SCHEDULE_HOST_CHECK", body)
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
		fmt.Println(string(e))
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

func ScheduleHostDowntime(data SchHostDowntime, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/SCHEDULE_HOST_DOWNTIME", body)
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
		fmt.Println(string(e))
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

func ScheduleServiceCheck(data SchServiceCheck, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/SCHEDULE_SVC_CHECK", body)
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
		fmt.Println(string(e))
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

func ScheduleServiceDowntime(data SchServiceDowntime, server string, username string, password string, allowInsecure bool) (CommandResult, error) {

	var returnedBody CommandResult
	var returnedError APIResponseError
	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://"+server+"/api/command/SCHEDULE_SVC_DOWNTIME", body)
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
		fmt.Println(string(e))
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
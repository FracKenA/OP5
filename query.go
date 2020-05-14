package op5

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HostResult []struct {
	Host
}

func GetQuery(server Server, query string, columns string) ([]byte, string, int, error) {
	var returnedError APIResponseError

	customTransport := http.DefaultTransport.(*http.Transport).Clone()

	parameters := url.Values{}

	if query != "" {
		parameters.Add("query", query)
	}
	if columns != ""{
		parameters.Add("columns", columns)
	}

	myurl, err := url.Parse("https://" + server.Host + "/api/filter/query")
	if err != nil {
		fmt.Println(err)
	}

	myurl.RawQuery = parameters.Encode()



	fmt.Println(myurl)
	req, err := http.NewRequest("GET",myurl.String(), nil)
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

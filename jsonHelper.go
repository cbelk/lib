package http

import (
    "github.com/cbelk/lib/loggers"
	"encoding/json"
    "fmt"
    "net/http"
	"io/ioutil"
    "time"
)

// This function forms a HTTP request and returns the data (json) retrieved from the request.
func GetJSON(method, url, password, user, application, logfile string) (res *http.Response) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		loggers.Freakout(fmt.Sprintf("[-] %s: Error (%v) retrieving json. Terminating at %v.", application, err, time.Now()), logfile, application, err)
	}
	req.SetBasicAuth(user, password)
	client := http.Client{}
	res, err = client.Do(req)
	if err != nil {
		loggers.Freakout(fmt.Sprintf("[-] %s: Error (%v) retrieving json. Terminating at %v.", application, err, time.Now()), logfile, application, err)
	}
	return
}

// This function parses a JSON object and stores it in the datapackage passed in.
func ParseJSON(res *http.Response, datapackage interface{}, application, logfile string) (status int) {
	status = res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		loggers.Freakout(fmt.Sprintf("[-] %s: Error (%v) parsing orgs json. Terminating at %v.", application, err, time.Now()), logfile, application, err)
	}
	err = res.Body.Close()
	if err != nil {
		fmt.Printf("[-] %s: Error (%v) closing the HTTP response body. Not worthy of termination.", application, err)
	}
	if status == 200 {
		err = json.Unmarshal(body, &datapackage)
		if err != nil {
			loggers.Freakout(fmt.Sprintf("[-] %s: Error (%v) parsing orgs json. Terminating at %v.", application, err, time.Now()), logfile, application, err)
		}
	}
	return
}

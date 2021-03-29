package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Run() error {

	// create web client
	//resp, err := http.Get("https://content.guardianapis.com/sections?api-key=43157614-9bc8-4fbe-985c-0c9e315c0c34")
	//if err != nil {
	//	return err
	//}

	client := http.DefaultClient

	req, _ := http.NewRequest(http.MethodGet, "https://content.guardianapis.com/sections?api-key=43157614-9bc8-4fbe-985c-0c9e315c0c34", nil)
	req.Header["Authentication"] = []string{"Bearer APIKEY"}
	req.SetBasicAuth("uyser", "passwrod")
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("couldn't make request, got %v", resp.StatusCode)
	}

	//// self-signed certs? NO

	// build URL with args and API key
	//// auth headers? NO

	// make request

	// unmarshall response to object
	// build content for file

	var dat map[string]interface{}
	rawBody, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(rawBody, &dat); err != nil {
		return err
	}

	// create type for data
	// unmarshall into type
	// extract specifics
	// transform to []byte before writing
	//// b := []byte(str)
	//// b := json.Marshal(obj) // {"foo":"goo"}

	err := ioutil.WriteFile("/tmp/news", rawBody, 0644)
	if err != nil {
		return err
	}
	// write to file
	//// where does this save? -> /tmp/news.txt

	return nil
}

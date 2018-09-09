package runners

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/codecyclist/kp-cli/cmd"
	"github.com/codecyclist/kp-cli/domain"
)

// PublishPaste uploads a Paste to the given Koalapaste Instance
func PublishPaste(appConfig cmd.AppConfig, paste domain.Paste) {

	jsonData, err := json.Marshal(paste)
	var apiEndpoint = appConfig.Stage + "/api/v1/pastes/"

	resp, err := http.Post(apiEndpoint, "application/json", strings.NewReader(string(jsonData)))

	if err == nil {
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			var paste domain.Paste
			unmarshalError := json.Unmarshal(bodyBytes, &paste)
			if unmarshalError == nil {
				pasteLink := appConfig.Stage + "/pastes/" + paste.Id
				fmt.Println(pasteLink)
			} else {
				fmt.Println("Unable to read response")
				fmt.Println(unmarshalError)
			}
		}
	} else {
		fmt.Println("unable to send paste")
	}
}

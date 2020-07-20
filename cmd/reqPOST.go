package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func reqPOST(endpoint string, bodyPOST string) {

	body := strings.NewReader(bodyPOST)
	req, err := http.NewRequest("POST", viper.GetString("baseURL")+endpoint, body)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Private-Token", viper.GetString("PersonalAccessToken"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
}

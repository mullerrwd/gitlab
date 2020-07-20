package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func reqGet(endpoint string, target interface{}) {
	req, err := http.NewRequest("GET", viper.GetString("baseURL")+endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Private-Token", viper.GetString("PersonalAccessToken"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		fmt.Println(err)
	}
}

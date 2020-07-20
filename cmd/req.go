package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func req(reqType string, endpoint string) {
	req, err := http.NewRequest(reqType, viper.GetString("baseURL")+endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Private-Token", viper.GetString("PersonalAccessToken"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
}

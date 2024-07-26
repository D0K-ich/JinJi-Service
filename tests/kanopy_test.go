package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"


	"net/http"
	"testing"

)


func TestShikiApi(t *testing.T) { //get auth code
	url := "https://shikimori.one/oauth/token"

	var data, _ = json.Marshal(&Data{
		GrantType:    "authorization_code",
		ClientId:     "dG-JqoHk3UVI5l-7vE9EvHGm2TLokE6RLLrUlzPXX4w",
		ClientSecret: "rw93sbyIN0gjvys0Fo96YPWoaIQY5t6MkDR9zPeN0ao",
		Code:         "gx_ctwbiM8klgQtMh8NAybNlYtZpdW_saF5YPYNhw_E",
		RedirectUri:  "urn:ietf:wg:oauth:2.0:oob",
	})

	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Kanopy")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

type Data struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
}
package gpt

import (
	"bytes"
	"encoding/json"
	"github.com/kr/pretty"
	"io"
	"net/http"
)

func(d *DocGPT) GetChats() {
	var err error
	var request []byte
	if request, err = json.Marshal(map[string]interface{}{"offset":2}); err != nil {return}

	var req *http.Request
	if req, err = http.NewRequest("POST", "http://localhost:3001/api/v1/admin/workspace-chats", bytes.NewReader(request)); err != nil {return}
	req.Header.Set("Authorization", "Bearer MEH45JA-48AMFV0-K4BMSBS-VT239C6")

	var response *http.Response
	var client = &http.Client{}
	if response, err = client.Do(req); err != nil {return}

	defer response.Body.Close()

	var response_bytes []byte
	if response_bytes, err = io.ReadAll(response.Body); err != nil {return}

	//var answer_gpt *ResponseGPT
	//if err = json.Unmarshal(response_bytes, &answer_gpt); err != nil {return}

	pretty.Println(string(response_bytes))
}

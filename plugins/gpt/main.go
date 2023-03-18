package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wa/api"
)

func Chat3(message map[string]string) string {
	url := "https://experimental.willow.vectara.io/v1/chat/completions"

	api.Messages = append(api.Messages, message)

	api.Counter++
	fmt.Println("A call recieved by Chat3")

	// Create the request data
	data := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    api.Messages,
		"max_tokens":  1500,
		"temperature": 0,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ready to call the API")
	// Create the request
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("customer-id", "2850595798")
	req.Header.Add("x-api-key", "zqt_qeij1u1-_mMgdjZqzf8IkQlmEQDLsZXDtKq52Q")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	text := string(body)
	fmt.Println(text)

	var result map[string]interface{}
	json.Unmarshal([]byte(text), &result)
	var answer string

	fmt.Println("answer:", answer)

	if result["choices"] != nil {
		x := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"]
		if x != nil {
			str, ok := x.(string)
			if ok {
				answer = strings.TrimLeft(str, "\n") // remove the leading \n
				fmt.Println("answer 2:", answer)
			} else {
				answer = "Sorry, try again"
				fmt.Println("x is not a string")
			}
		}
	}
	fmt.Println("answer3:", answer)
	return answer
}

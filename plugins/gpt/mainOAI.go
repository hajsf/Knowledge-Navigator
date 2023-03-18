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

func ChatGPT3(message map[string]string) string {
	url := "https://api.openai.com/v1/chat/completions"
	API_KEY := ""

	api.Messages = append(api.Messages, message)

	api.Counter++
	fmt.Println("A call recieved by ChatGPT3")
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+API_KEY)

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

	fmt.Println(answer)

	if result["choices"] != nil {
		x := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"]
		if x != nil {
			str, ok := x.(string)
			if ok {
				answer = strings.TrimLeft(str, "\n")
				fmt.Print(strings.TrimLeft(str, "\n")) // remove the leading \n
			} else {
				answer = "Sorry, try again"
				fmt.Println("x is not a string")
			}
		}
	}
	return answer
}
